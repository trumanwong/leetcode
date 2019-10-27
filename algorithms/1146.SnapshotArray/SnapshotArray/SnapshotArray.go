package SnapshotArray

type SnapshotArray struct {
	Nums map[int]map[int]int
	SnapId int
}


func Constructor(length int) SnapshotArray {
	nums := make(map[int]map[int]int)
	snapId := 0
	for i := 0; i < length; i++ {
		temp := make(map[int]int)
		temp[snapId] = 0
		nums[i] = temp
	}
	arr := SnapshotArray{Nums: nums, SnapId: snapId}
	return arr
}


func (this *SnapshotArray) Set(index int, val int)  {
	this.Nums[index][this.SnapId] = val
}


func (this *SnapshotArray) Snap() int {
	this.SnapId++
	return this.SnapId - 1
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
	for i := snap_id; i >= 0; i-- {
		if v, ok := this.Nums[index][i]; ok {
			return v
		}
	}
	return -1
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */