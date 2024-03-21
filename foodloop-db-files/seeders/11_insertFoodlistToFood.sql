INSERT INTO foodloop.foodlistToFood (foodlistID, foodID, foodIndex)
SELECT fl.foodlistID, ft.foodID,
       ROW_NUMBER() OVER (PARTITION BY fl.foodlistID ORDER BY ft.foodID) AS foodIndex
FROM foodloop.foodlist fl
         JOIN foodloop.tag t ON fl.foodlistName = t.tagName
         JOIN foodloop.foodToTag ft ON t.tagID = ft.tagID;