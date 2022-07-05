/*
 * @Description: 请输入....
 * @Author: Gavin
 * @Date: 2022-07-04 17:52:44
 * @LastEditTime: 2022-07-04 19:15:23
 * @LastEditors: Gavin
 */
package interface_demo

type USBER interface {
	open()
	close()
}

type Door struct {
	name string
}

func (d Door) open() {

}
func (d Door) close() {}

func Run() {
	//实现一个接口
	var c USBER
	c = Door{
		name: "20",
	}
	c.close()
}
