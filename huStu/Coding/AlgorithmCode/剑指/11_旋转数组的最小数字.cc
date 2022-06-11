class Solution {
public:
    int minArray(vector<int>& numbers) {
        int left = 0;
        int right = numbers.size()-1;
        while(left < right)
        {
            int mid = (right-left)/2+left;
            if(numbers[mid] == numbers[right])
            {
                right--;
            }
            else if(numbers[mid] < numbers[right])
            {
                right = mid;
            }
            else
            {
                left = mid+1;
            }
        }
        return numbers[right];
    }
};