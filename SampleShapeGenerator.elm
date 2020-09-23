-- Cerner_2^5_2020
-- Create pictures from simple shapes. Like a tree!

import Playground exposing (..)


main =
  picture
    [ rectangle brown 40 200
        |> moveDown 80
    , circle green 100
        |> moveUp 100
    ]