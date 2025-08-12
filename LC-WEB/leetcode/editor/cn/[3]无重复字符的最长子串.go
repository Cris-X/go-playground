//给定一个字符串 s ，请你找出其中不含有重复字符的 最长 子串 的长度。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "abcabcbb"
//输出: 3 
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 
//
// 示例 2: 
//
// 
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 
//
// 示例 3: 
//
// 
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 
//
// 
//
// 提示： 
//
// 
// 0 <= s.length <= 5 * 10⁴ 
// s 由英文字母、数字、符号和空格组成 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 10975 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) (res int) {
	charIndex := make(map[byte]int)
	left := 0
	for right := 0; right < len(s); right ++ {
		// charIndex[字符] = 该字符最新的位置
		//字符串: "abba"
		//索引:    0123
		//
		//right=0: 'a' → charIndex['a']=0
		//right=1: 'b' → charIndex['b']=1
		//right=2: 'b' → 发现重复！
		//charIndex['b']=1 → left跳到2（1+1）
		//一步到位，跳过了'a'和第一个'b'
		//right=3: 'a' → charIndex['a']=0，但0 < left(2)
		//说明这个'a'在窗口外，不用管

		// 如果字符已存在，更新左边界
		if lastPos, exists := charIndex[s[right]]; exists && lastPos >= left {
			left = lastPos + 1
		}

		// 更新字符位置
		charIndex[s[right]] = right

		res = max(res, right-left+1)
	}

	return
}
//leetcode submit region end(Prohibit modification and deletion)
