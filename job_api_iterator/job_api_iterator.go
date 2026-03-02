package jobapiiterator

import (
	"fmt"
	"strconv"
)

// 具象iterator。求人データをiteratorとして処理する
type JobAPIIterator struct{
	CurrentPage int
	Buffer []Job //APIから取得した1ページ分のデータを一時保存
	Index int //現在のバッファ内のどこをみているか
	IsDone bool //全データ取り切ったかどうか
}

type Job struct{
	ID int
	Title string
}

// ダミーのAPI呼び出し関数。今回は１ページに２件返す設定
func (it *JobAPIIterator) FetchJobsFromAPI(page int)[]Job{
	if page > 3{
		return []Job{} // 3ページ目で終わり
	}
	fmt.Printf("\n [API 通信発生] ページ %d を取得中... \n", page)
	return []Job{
		{ID: page*10 + 1, Title: "エンジニア" + strconv.Itoa(page*10+1)},
		{ID: page*10 + 2, Title: "デザイナー" + strconv.Itoa(page*10+2)},
	}
}

// バッファが空または全て読み切った場合、次のページを裏側でフェッチする
func (it *JobAPIIterator) HasNext() bool{
	// バッファが空、または全て読み切った場合、次のページを裏側でフェッチする
	if it.Index >= len(it.Buffer) && !it.IsDone{
		it.CurrentPage++
		it.Buffer = it.FetchJobsFromAPI(it.CurrentPage)
		it.Index = 0

		if len(it.Buffer) == 0{
			it.IsDone = true
		}
	}

	return !it.IsDone
}

// バッファから1件取り出してインデックスを１つ進める
func (it *JobAPIIterator) Next() Job{
	// バッファから1件取り出して、インデックスを進める
	job := it.Buffer[it.Index]
	it.Index++
	return job
}