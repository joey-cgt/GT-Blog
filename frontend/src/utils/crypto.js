/**
 * 密码加密与解密工具
 * 注意：此实现仅为示例，在实际生产环境中，建议使用更复杂的加密方案
 */

/**
 * 生成随机密钥
 * @returns {string} 随机生成的密钥
 */
function generateKey() {
  const array = new Uint8Array(16);
  window.crypto.getRandomValues(array);
  return Array.from(array, byte => byte.toString(16).padStart(2, '0')).join('');
}

/**
 * 获取或创建加密密钥
 * @returns {string} 加密密钥
 */
export function getEncryptionKey() {
  let key = localStorage.getItem('encryptionKey');
  if (!key) {
    key = generateKey();
    localStorage.setItem('encryptionKey', key);
  }
  return key;
}

/**
 * 简单的加密函数（生产环境请使用更安全的加密方式）
 * @param {string} text - 要加密的文本
 * @param {string} key - 加密密钥
 * @returns {string} 加密后的文本
 */
export function encrypt(text, key) {
  try {
    const encoder = new TextEncoder();
    const data = encoder.encode(text);
    
    // 简单的异或加密（仅用于演示，生产环境应使用AES等标准加密算法）
    const encryptedData = new Uint8Array(data.length);
    for (let i = 0; i < data.length; i++) {
      encryptedData[i] = data[i] ^ key.charCodeAt(i % key.length);
    }
    
    // 将加密后的数据转换为Base64
    return btoa(String.fromCharCode(...encryptedData));
  } catch (error) {
    console.error('加密失败:', error);
    return '';
  }
}

/**
 * 解密函数
 * @param {string} encryptedText - 加密后的文本
 * @param {string} key - 解密密钥
 * @returns {string} 解密后的文本
 */
export function decrypt(encryptedText, key) {
  try {
    // 从Base64解码
    const encryptedData = new Uint8Array(atob(encryptedText).split('').map(char => char.charCodeAt(0)));
    
    // 简单的异或解密
    const decryptedData = new Uint8Array(encryptedData.length);
    for (let i = 0; i < encryptedData.length; i++) {
      decryptedData[i] = encryptedData[i] ^ key.charCodeAt(i % key.length);
    }
    
    const decoder = new TextDecoder();
    return decoder.decode(decryptedData);
  } catch (error) {
    console.error('解密失败:', error);
    return '';
  }
}

/**
 * 安全地保存密码到localStorage
 * @param {string} password - 要保存的密码
 */
export function savePassword(password) {
  const key = getEncryptionKey();
  const encryptedPassword = encrypt(password, key);
  localStorage.setItem('rememberedPassword', encryptedPassword);
}

/**
 * 从localStorage中安全地获取密码
 * @returns {string|null} 解密后的密码，如果不存在则返回null
 */
export function getPassword() {
  const encryptedPassword = localStorage.getItem('rememberedPassword');
  if (!encryptedPassword) {
    return null;
  }
  
  const key = getEncryptionKey();
  const password = decrypt(encryptedPassword, key);
  return password || null;
}

/**
 * 删除保存的密码
 */
export function removePassword() {
  localStorage.removeItem('rememberedPassword');
}