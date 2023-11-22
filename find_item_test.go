package main

// func TestFindItem(t *testing.T) {
// 	db := Open()

// 	defer Close(db)

// 	t.Run("when Item is not exist", func(t *testing.T) {
// 		res := FindItem(db, 1754)

// 		if nil != res {
// 			t.Errorf(M, nil, res)
// 		}
// 	})

// 	t.Run("when Item is exist", func(t *testing.T) {
// 		params := &ItemParams {
// 			Date: date.New(2023, 11, 21),
// 			Formula: "2.1 + 1.8",
// 			CategoryID: 11,
// 		}

// 		item := NewItem(db)

// 		item.Create(params)

// 		fmt.Println(item.ID)

// 		// i2 := FindItem(db, item.ID)

// 		// fmt.Println(i2)
// 	})
// }
