package jobapiiterator

import (
	"fmt"
	"iter"
	"strconv"
)

type Job struct{
	ID int
	Title string
}

// ダミーのAPI呼び出し関数。今回は１ページに２件返す設定
func FetchJobsFromAPI(page int)[]Job{
	if page > 3{
		return []Job{} // 3ページ目で終わり
	}
	fmt.Printf("\n [API 通信発生] ページ %d を取得中... \n", page)
	return []Job{
		{ID: page*10 + 1, Title: "エンジニア" + strconv.Itoa(page*10+1)},
		{ID: page*10 + 2, Title: "デザイナー" + strconv.Itoa(page*10+2)},
	}
}


func FetchAllJobs() iter.Seq[Job] {
	// 戻り値は関数→どんな関数？：yieldという関数を受け取る
	return func(yield func(Job) bool){
		page := 1 //状態は単なるローカル変数

		for {
			// APIから1ページ分を取得
			jobs := FetchJobsFromAPI(page)
			if len(jobs) == 0{
				return
			}
			// 取得したデータを1件ずつ、使う側(for range)に「yield(渡す)」する
			for _, job := range jobs{
				// yield(job)を呼ぶと、main側のforループの中身が実行される
				// もしmain側でbreakされたら、yieldはfalseを返す。これによって、クライアンが処理をやめた瞬間にクリーンアップを実行することが自然にかける
				if !yield(job){
					fmt.Println(" [System] クライアントがループを中断しました。通信を遮断します.")
					return //即座に終了してリソースを解放
				}
			}
			page++
		}
	}
}
