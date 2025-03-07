package timer

import (
	"gonet/base"
	"sync"
	"sync/atomic"
	"time"
)

const (
	TIME_NEAR_SHIFT  = 8
	TIME_NEAR        = (1 << TIME_NEAR_SHIFT)
	TIME_LEVEL_SHIFT = 6
	TIME_LEVEL       = (1 << TIME_LEVEL_SHIFT)
	TIME_NEAR_MASK   = (TIME_NEAR - 1)
	TIME_LEVEL_MASK  = (TIME_LEVEL - 1)
	TICK_INTERVAL    = 10 * time.Millisecond
)

//先搞清楚下面的单位
//1秒=1000毫秒 milliseconds
//1毫秒=1000微秒 microseconds
//1微秒=1000纳秒 nanoseconds
//整个timer中毫秒的精度都是10ms，
//也就是说毫秒的一个三个位，但是最小的位被丢弃
type (
	TimerHandle func()
	TimerNode   struct {
		next   *TimerNode
		expire uint32
		handle TimerHandle
		id     *int64
		//readd
		time  uint32
		bOnce bool
	}

	//这个队列可以换成无锁队列
	LinkList struct {
		head TimerNode
		tail *TimerNode
	}

	Timer struct {
		near          [TIME_NEAR]LinkList     //临近的定时器数组
		t             [4][TIME_LEVEL]LinkList //四个级别的定时器数组
		lock          sync.Mutex              //锁
		time          uint32                  //计数器
		starttime     uint32                  //程序启动的时间点，timestamp，秒数
		current       uint64                  //从程序启动到现在的耗时，精度10毫秒级
		current_point uint64                  //当前时间，精度10毫秒级
		pTimer        *time.Ticker            //定时器
		loop_node     []*TimerNode
	}

	Op struct {
		bOnce bool
	}

	OpOption func(*Op)
)

var (
	TIMER *Timer
	g_Id  int64
)

func (this *TimerNode) LoadId() int64 {
	return atomic.LoadInt64(this.id)
}

func (op *Op) applyOpts(opts []OpOption) {
	for _, opt := range opts {
		opt(op)
	}
}

func WithOnce() OpOption {
	return func(op *Op) {
		op.bOnce = true
	}
}

func init() {
	TIMER = &Timer{}
	TIMER.Init()
}

func uuid() int64 {
	return atomic.AddInt64(&g_Id, 1)
}

//清空链表，返回链表第一个结点
func linkClear(list *LinkList) *TimerNode {
	ret := list.head.next
	list.head.next = nil
	list.tail = &list.head
	return ret
}

//将结点放入链表
func link(list *LinkList, node *TimerNode) {
	list.tail.next = node
	list.tail = node
	node.next = nil
}

//创建一个定时器
func (this *Timer) Init() {
	for i := 0; i < TIME_NEAR; i++ {
		linkClear(&this.near[i])
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < TIME_LEVEL; j++ {
			linkClear(&this.t[i][j])
		}
	}

	this.current = 0
	this.pTimer = time.NewTicker(TICK_INTERVAL)
	this.current_point = uint64(time.Now().UnixNano()) / uint64(TICK_INTERVAL)
	go this.run()
}

//添加一个定时器结点
func (this *Timer) addNode(node *TimerNode) {
	time := node.expire       //去看一下它是在哪赋值的
	current_time := this.time //当前计数
	//没有超时，或者说时间点特别近了
	if (time | TIME_NEAR_MASK) == (current_time | TIME_NEAR_MASK) {
		link(&this.near[time&TIME_NEAR_MASK], node)
	} else { //这里有一种特殊情况，就是当time溢出，回绕的时候
		i := 0
		mask := uint32(TIME_NEAR << TIME_LEVEL_SHIFT)
		for i = 0; i < 3; i++ { //看到i<3没，很重要很重要
			if (time | (mask - 1)) == (current_time | (mask - 1)) {
				break
			}
			mask <<= TIME_LEVEL_SHIFT //mask越来越大
		}
		//放入数组中
		link(&this.t[i][(time>>uint(TIME_NEAR_SHIFT+i*TIME_LEVEL_SHIFT))&TIME_LEVEL_MASK], node)
	}
}

//添加一个定时器
func (this *Timer) Add(id *int64, time uint32, handle TimerHandle, opts ...OpOption) *TimerNode {
	op := Op{}
	op.applyOpts(opts)
	node := &TimerNode{expire: time + this.time, handle: handle,
		time: time, bOnce: op.bOnce, id: id} //超时时间+当前计数
	this.lock.Lock()
	this.addNode(node)
	this.lock.Unlock()
	return node
}

//删除一个定时器
func (this *Timer) Delete(id *int64) {
	atomic.StoreInt64(id, 0)
}

//移动某个级别的链表内容
func (this *Timer) moveList(level int, idx int) {
	current := linkClear(&this.t[level][idx])
	for current != nil {
		temp := current.next
		this.addNode(current)
		current = temp
	}
}

//这是一个非常重要的函数
//定时器的移动都在这里
func (this *Timer) shift() {
	mask := uint32(TIME_NEAR)
	this.time += 1
	ct := this.time
	if ct == 0 { //time溢出了
		this.moveList(3, 0) //这里就是那个很重要的3
	} else { //time正常
		time := ct >> TIME_NEAR_SHIFT
		i := 0

		for (ct & (mask - 1)) == 0 {
			idx := time & TIME_LEVEL_MASK
			if idx != 0 {
				this.moveList(i, int(idx))
				break
			}
			mask <<= TIME_LEVEL_SHIFT //mask越来越大
			time >>= TIME_LEVEL_SHIFT //time越来越小
			i += 1
		}
	}
}

//派发消息到目标服务消息队列
func (this *Timer) dispatch(current *TimerNode) {
	for current != nil {
		id := current.LoadId()
		if id != 0 {
			current.handle()
			if !current.bOnce {
				this.loop_node = append(this.loop_node, current)
			}
		}
		current = current.next
	}
}

//派发消息
func (this *Timer) execute() {
	idx := this.time & TIME_NEAR_MASK

	for this.near[idx].head.next != nil {
		current := linkClear(&this.near[idx])
		this.lock.Unlock()
		// dispatch don't need lock T
		this.dispatch(current)
		this.lock.Lock()
		for _, v := range this.loop_node {
			v.expire = v.time + this.time
			this.addNode(v)
		}
		this.loop_node = []*TimerNode{}
	}
}

//时间更新好了以后，这里检查调用各个定时器
func (this *Timer) advace() {
	this.lock.Lock()
	// try to dispatch timeout 0 (rare condition)
	this.execute()
	// shift time first, and then dispatch timer message
	this.shift()
	this.execute()
	this.lock.Unlock()
}

//在线程中不断被调用
//调用时间 间隔为微秒
func (this *Timer) update() {
	cp := uint64(time.Now().UnixNano()) / uint64(TICK_INTERVAL)
	if cp < this.current_point {
		this.current_point = cp
	} else if cp != this.current_point {
		diff := cp - this.current_point
		this.current_point = cp //当前时间，毫秒级
		this.current += diff    //从启动到现在耗时
		for i := uint64(0); i < diff; i++ {
			this.advace() //注意这里
		}
	}
}

func (this *Timer) loop() bool {
	defer func() {
		if err := recover(); err != nil {
			base.TraceCode(err)
		}
	}()

	select {
	case <-this.pTimer.C:
		this.update()
	}
	return true
}

func (this *Timer) run() {
	for {
		if !this.loop() {
			break
		}
	}
	this.pTimer.Stop()
}

func RegisterTimer(id *int64, duration time.Duration, handle TimerHandle, opts ...OpOption) {
	TIMER.Add(id, uint32(duration/TICK_INTERVAL), handle, opts...)
}

func StopTimer(id *int64) {
	if id != nil {
		TIMER.Delete(id)
	}
}
