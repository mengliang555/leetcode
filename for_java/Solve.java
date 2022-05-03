import java.util.HashMap;
import java.util.Map;

public class Solve {
    // 两数之和 https://leetcode-cn.com/problems/two-sum/
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> store = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (store.containsKey(target - nums[i])) {
                return new int[]{store.get(target - nums[i]), i};
            } else {
                store.put(nums[i], i);
            }
        }
        return new int[]{-1, -1};
    }
}