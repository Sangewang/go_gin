/* 是否是公司邮箱*/
export function isByteDanceEmail(str) {
  const reg = /^[a-z0-9](?:[-_.+]?[a-z0-9]+)*@bytedance\.com$/i;
  return reg.test(str.trim());
}