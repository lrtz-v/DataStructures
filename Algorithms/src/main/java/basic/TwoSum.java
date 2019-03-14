package basic;
import java.util.HashMap;
import java.util.ArrayList;
import java.util.List;


public class TwoSum {

    public static List<Integer> twoSum(List<Integer> nums, Integer target) {
        HashMap<Integer, Integer> numHash = new HashMap<Integer, Integer>();
        List<Integer> l = new ArrayList<Integer>();
        for (int i = 0; i< nums.size(); i++) {
            numHash.put(nums.get(i), i);
        }
        for (int i = 0; i<nums.size(); i++) {
            int tmpVal = target - nums.get(i);
            Integer value = numHash.get(tmpVal);
            if ( value != null && i != value ) {
                l.add(i);
                l.add(value);
                break;
            }
        }
        return l;
    }
}
