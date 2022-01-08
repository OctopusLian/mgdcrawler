/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-01-08 13:09:10
 * @LastEditors: neozhang
 * @LastEditTime: 2022-01-08 13:09:11
 */
package model

type Car struct {
	Name         string
	Price        float64
	ImageURL     string
	Size         string
	Fuel         float64
	Transmission string
	Engine       string
	Displacement float64 // 排量
	MaxSpeed     float64
	Acceleration float64
}
