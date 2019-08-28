<?php

class Solution
{
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    public function twoSum($nums, $target)
    {
        $result = [];
        $count = count($nums);
        for ($i = 0; $i < $count - 1; $i++) {
            for ($j = $i + 1; $j < $count; $j++) {
                if ($nums[$i] + $nums[$j] == $target) {
                    $result = [
                        $i, $j,
                    ];
                    break;
                }
            }
        }
        return $result;

    }
}

$nums = [3, 2, 4];
$target = 6;
$solution = new Solution;
$result = $solution->twoSum($nums, $target);
var_dump($result);
