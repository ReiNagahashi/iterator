package main

import (
	"fmt"
	"iterator/iterator"
	jobapiiterator "iterator/job_api_iterator"
)

// 目的：外部APIから大量のデータを取得して処理する
func main(){
	var iterator iterator.Iterator[jobapiiterator.Job] = &jobapiiterator.JobAPIIterator{
		CurrentPage: 0,
		Index: 0,
	}

	fmt.Println("=== 求人データ処理開始 ===")

	count := 0
	for iterator.HasNext(){
		job := iterator.Next()
		fmt.Printf("処理中: ID=%d, Title=%s\n", job.ID, job.Title)
		count++
	}

	fmt.Printf("処理回数：%d", count)
}