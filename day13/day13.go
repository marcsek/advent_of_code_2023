//package main
//
//import (
//	"fmt"
//	"os"
//	"slices"
//	"strings"
//)
//
//func main() {
//	input := strings.Split(readFile("sample2.txt"), "\n")
//
//	total := 0
//	for _, ref := range makeIntoReflec(input[:len(input)-1]) {
//		//printMaze(rotateSlice(ref[:len(ref)-1]))
//		if cl, cr := tryPatter(ref[:len(ref)-1]); cl != -1 {
//			total += cr
//			fmt.Println(cl, cr)
//		} else {
//			cl, cr := tryPatter(rotateSlice(ref[:len(ref)-1]))
//			total += cr * 100
//			fmt.Println(cl, cr)
//		}
//	}
//	fmt.Println(total)
//}
//
//func makeIntoReflec(input []string) [][]string {
//	res := [][]string{}
//	newRef := []string{}
//	for _, r := range input {
//		if r != "" {
//			newRef = append(newRef, r)
//		} else {
//			res = append(res, newRef)
//			newRef = []string{}
//		}
//	}
//	res = append(res, newRef)
//	return res
//}
//
//func tryPatter(patter []string) (int, int) {
//	cand := [][]int{}
//
//	for i, r := range patter {
//		fnd1 := tryRow(r, true)
//		fnd2 := tryRow(r, false)
//
//		cm := true
//		s1 := fnd2
//		s2 := fnd1
//
//		if len(fnd1) > len(fnd2) {
//			cm = false
//			s1 = fnd1
//			s2 = fnd2
//		}
//
//		for off, f1 := range s1 {
//			if off+1 >= len(s2) {
//				break
//			}
//			for _, f2 := range s2[off+1:] {
//				//fmt.Println(f1, f2, i)
//				//l1, r1 := fnd1[0][0], fnd1[0][1]
//				//l2, r2 := fnd2[0][0], fnd2[0][1]
//				l1, r1 := f1[0], f1[1]
//				l2, r2 := f2[0], f2[1]
//				if cm {
//					l1, r1 = f2[0], f2[1]
//					l2, r2 = f1[0], f1[1]
//				}
//				//fmt.Println(l1, r1, l2, r2)
//
//				if i != 0 {
//					nCnd := [][]int{}
//					for _, cnd := range cand {
//						cl, cr := cnd[0], cnd[1]
//						if (cl == l1 && cr == r1) || (cl == l2 && cr == r2) {
//							nCnd = append(nCnd, cnd)
//						}
//					}
//					//fmt.Println(nCnd, l1, r1, l2, r2, i)
//					if l1 == l2 && r1 == r2 {
//						cand = append(cand, nCnd...)
//					} else {
//						cand = slices.Clone(nCnd)
//					}
//				} else {
//					if l1 != l2 && r1 != r2 {
//						cand = append(cand, []int{l1, r1})
//						cand = append(cand, []int{l2, r2})
//					} else {
//						cand = append(cand, []int{l1, r1})
//					}
//				}
//			}
//		}
//	}
//
//	if len(cand) == 0 {
//		return -1, -1
//	}
//	return cand[0][0], cand[0][1]
//}
//
//func tryRow(row string, rToL bool) [][]int {
//	lPtr := 0
//	rPtr := len(row) - 1
//	found := [][]int{}
//	for lPtr < rPtr {
//		if row[lPtr] == row[rPtr] {
//			in1, in2 := inside(row, lPtr, rPtr)
//			if in1 != -1 {
//				fmt.Println(in1, in2)
//				found = append(found, []int{in1, in2})
//			}
//		}
//		if rToL {
//			lPtr++
//		} else {
//			rPtr--
//		}
//	}
//	if len(found) == 0 {
//		return [][]int{{-1, -1}}
//	}
//	fmt.Println(found)
//	return found
//}
//
//func inside(row string, lPtr int, rPtr int) (int, int) {
//	for lPtr < rPtr {
//		if row[lPtr] != row[rPtr] {
//			return -1, -1
//		}
//		lPtr++
//		rPtr--
//	}
//	if rPtr == lPtr {
//		return -1, -1
//	}
//
//	return lPtr - 1, rPtr + 1
//}
//
//func rotateSlice(slice []string) []string {
//	rotated := []string{}
//
//	for x := range slice[0] {
//		row := ""
//		for y := range slice {
//			row += string(slice[y][x])
//		}
//		rotated = append(rotated, row)
//	}
//
//	return rotated
//}
//
//func printMaze(maze []string) {
//	for _, r := range maze {
//		fmt.Println(r)
//	}
//}
//
//func readFile(fn string) string {
//	data, err := os.ReadFile(fn)
//
//	if err != nil {
//		panic("Couldn't read input file")
//	}
//
//	return string(data)
//}
