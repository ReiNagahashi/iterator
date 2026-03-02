package main

import (
	"fmt"
	jobapiiterator "iterator/job_api_iterator"
)

// 目的：外部APIから大量のデータを取得して処理する
func main(){

	fmt.Println("=== 求人データ処理開始 ===")

	count := 0
	for job := range jobapiiterator.FetchAllJobs(){
		fmt.Printf("処理中: ID=%d, Title=%s\n", job.ID, job.Title)
		count++

		// 3件処理したら途中でやめてみる 
		if count == 3{
			fmt.Println(">> 3件処理したのでbreakします")
			break
		}
	}

	fmt.Println("=== 処理完了 ===")
}