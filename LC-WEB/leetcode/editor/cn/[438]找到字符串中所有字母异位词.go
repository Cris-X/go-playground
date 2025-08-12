//给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。 
//
// 
//
// 示例 1: 
//
// 
//输入: s = "cbaebabacd", p = "abc"
//输出: [0,6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
// 
//
// 示例 2: 
//
// 
//输入: s = "abab", p = "ab"
//输出: [0,1,2]
//解释:
//起始索引等于 0 的子串是 "ab", 它是 "ab" 的异位词。
//起始索引等于 1 的子串是 "ba", 它是 "ab" 的异位词。
//起始索引等于 2 的子串是 "ab", 它是 "ab" 的异位词。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= s.length, p.length <= 3 * 10⁴ 
// s 和 p 仅包含小写字母 
// 
//
// Related Topics 哈希表 字符串 滑动窗口 👍 1731 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s, p string) (ans []int) {
	cntP := [26]int{} // 统计 p 的每种字母的出现次数
	cntS := [26]int{} // 统计 s 的长为 len(p) 的子串 s' 的每种字母的出现次数
	for _, c := range p {
		cntP[c-'a']++ // 统计 p 的字母
	}
	for right, c := range s {
		cntS[c-'a']++ // 右端点字母进入窗口
		left := right - len(p) + 1
		if left < 0 { // 窗口长度不足 len(p)
			continue
		}
		if cntS == cntP { // s' 和 p 的每种字母的出现次数都相同
			ans = append(ans, left) // s' 左端点下标加入答案
		}
		cntS[s[left]-'a']-- // 左端点字母离开窗口
	}
	return
}
//leetcode submit region end(Prohibit modification and deletion)
