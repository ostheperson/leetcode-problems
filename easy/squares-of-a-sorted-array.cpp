vector<int> sortedSquares(vector<int>& nums) {
    /*
    [-5,-1,0,3,4]
    [16, 25]
    i = 0
    j = 4
    k = 4

    if
    */
    int n = nums.size();
    int i =  0, j = n - 1;
    vector<int> res = vector<int>(n,-1);
    for(int k=n-1;k>=0;--k){
        if(abs(nums[i])>=abs(nums[j])){
            res[k] = nums[i]*nums[i];
            i++;
        }else{
            res[k] = nums[j]*nums[j];
            j--;
        }
    }
    return res;
}