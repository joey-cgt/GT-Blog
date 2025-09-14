<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { articles } from '../store/blog.js'
import Comments from './Comments.vue'
import MarkdownRenderer from './MarkdownRenderer.vue'

const route = useRoute()
const articleId = computed(() => parseInt(route.params.id))
const article = computed(() => {
  return articles.find(a => a.id === articleId.value) || null
})

// 模拟文章内容（实际项目中可能从API获取）
const content = ref('')

// 点赞功能
const likeCount = ref(0)
const handleLike = () => {
  likeCount.value++
  // 这里可以添加API调用保存点赞数据
  console.log('点赞成功', likeCount.value)
}

// 滚动到评论区
const scrollToComments = () => {
  const commentsSection = document.getElementById('comments')
  if (commentsSection) {
    commentsSection.scrollIntoView({ behavior: 'smooth' })
  } else {
    console.log('评论区未找到')
  }
}

onMounted(() => {
  // 模拟文章内容（实际项目中应该从API获取）
  if (article.value) {
    content.value = generateDummyContent(article.value.title)
  }
})

// 生成模拟Markdown内容的辅助函数
function generateDummyContent(title) {
  return `
# Web3：下一代互联网的范式革命
- 选项一
  - 选项二

- 选项二

- 选项三

\`\`\`java
public static void main(String[] args) {
  System.out.println("Hello, World!");
}
\`\`\`

## 引言：从"只读"到"读写"，再到"拥有"

互联网的演进历程，如同一部波澜壮阔的技术史诗。我们经历了 **Web 1.0**，一个"只读"（Read-Only）的时代。那时的互联网是静态的，用户主要是被动的信息接收者，浏览着由少数公司或个人创建的网页门户（如雅虎、新浪）。随后，我们进入了 **Web 2.0**，一个"读写"（Read-Write）的时代。社交媒体（如Facebook、Twitter、微博）、用户生成内容（UGC）平台（如YouTube、抖音）和协作工具（如Wikipedia）的兴起，使得用户从消费者转变为内容的创造者和传播者。互联网变得前所未有的动态、社交化和中心化。

然而，Web 2.0 的繁荣背后隐藏着巨大的代价：**数据的中心化垄断和价值的单向流动**。少数科技巨头（如Google, Meta, Amazon）通过提供"免费"的服务，收集了海量的用户数据，并以此构建了其商业帝国的护城河。用户创造了绝大部分价值（内容、关系、注意力），但却无法真正拥有自己的数字身份和数据，更无法分享平台成长所带来的巨大收益。数据泄露、隐私滥用、算法操纵、平台封禁等问题日益凸显。

正是在这样的背景下，**Web3** 应运而生。它被广泛誉为互联网的下一个阶段，一个"拥有"（Own）的时代。Web3 的核心愿景是构建一个**去中心化的、用户拥有其数字身份和数据、并通过代币经济（Token Economy）参与价值分配**的新型互联网。它并非要完全取代Web2.0，而是旨在通过一套全新的技术栈和经济模型，重塑数字世界的权力结构和生产关系。

本文将深入探讨Web3的核心理念、关键技术、生态构成、面临的挑战以及未来展望，为您全面解析这一正在发生的范式革命。

---

## 第一章：Web3的核心理念与基本特征

Web3 并非单一技术的突进，而是一套融合了密码学、经济学和社会学的复杂思想体系。其核心可以概括为以下几个相互关联的原则：

### 1. 去中心化（Decentralization）
这是Web3最根本的特征。与Web2.0由中心化服务器和公司控制不同，Web3应用（称为DApps，去中心化应用）运行在由无数节点组成的分布式网络上（如区块链）。没有单一实体能够控制或关闭整个网络，确保了系统的抗审查性和韧性。权力从中心化的组织转移到分布式的社区。

### 2. 所有权（Ownership）
在Web3中，用户通过**加密密钥**和**智能合约**真正拥有自己的数字资产和数据。
*   **数字资产**：例如，你钱包里的比特币（BTC）、以太坊（ETH）或NFT（非同质化代币），完全由你的私钥控制。只要保管好私钥，没有任何第三方可以冻结或没收你的资产。
*   **数据**：你的身份信息、社交图谱、创作内容等可以存储在去中心化的存储网络中（如IPFS, Arweave），并由你授权给应用访问。应用为你服务，而非你为应用贡献数据。

### 3. 原生数字支付（Native Digital Payments）
Web3生态内置了强大的金融层，主要以**加密货币**（Cryptocurrency）的形式存在。这使得全球范围内的点对点（P2P）价值转移变得极其高效和低成本，无需依赖传统的银行或支付网关。这为微支付、机器经济（Machine-to-Machine payments）和全新的商业模式奠定了基础。

### 4. 信任最小化（Trust Minimization）与可验证性（Verifiability）
通过区块链的公开账本和密码学证明，Web3应用实现了"无需信任"（Trustless）——并非指不值得信任，而是指不需要信任某个中介或第三方。所有规则都编码在开源的智能合约中，所有交易记录和状态变更都是公开透明、可被任何人验证的。这极大地降低了合作和交易中的信任成本。

### 5. 社区治理（Community Governance）
许多Web3项目和协议都发行了**治理代币**（Governance Token）。持有这些代币的用户有权对项目的未来发展、资金使用等关键决策进行提案和投票。这试图将项目的控制权交给其用户和社区，实现"协议民主化"，形成一种去中心化自治组织（DAO）的形态。

---

## 第二章：Web3的技术基石

梦想需要技术来实现。Web3的宏伟愿景建立在以下几项突破性技术之上：

### 1. 区块链（Blockchain）
**区块链是Web3的底层结算层和信任基础**。它是一个按时间顺序排列的、由加密算法链接的"区块"构成的分布式账本。其核心特性包括：
*   **去中心化**：数据由网络中的多个节点共同维护，无需中心化机构。
*   **不可篡改性**：一旦数据被写入区块并得到确认，几乎不可能被更改。
*   **透明性**：所有交易记录对网络参与者公开可查。
*   **安全性**：通过工作量证明（PoW）、权益证明（PoS）等共识机制保障网络安全。

以太坊（Ethereum）是当前Web3生态中最主要的智能合约平台，但此外还有Solana、Avalanche、Polkadot等多个竞争性公链。

### 2. 智能合约（Smart Contracts）
**智能合约是运行在区块链上的自执行代码**。它定义了特定条件下自动执行的规则和逻辑。例如："如果A向地址X转入1个ETH，则自动将1个某NFT转入A的地址"。智能合约使得复杂的去中心化应用（如借贷、交易、游戏）成为可能，是Web3应用的"业务逻辑层"。

### 3. 加密货币与代币（Cryptocurrencies & Tokens）
*   **原生代币**：如BTC、ETH，是各自区块链网络的"燃料"，用于支付交易手续费和激励网络维护者。
*   **治理代币**：如UNI（Uniswap）、MKR（MakerDAO），赋予持有者参与协议治理的权利。
*   **NFT**：一种独特的、不可互换的数字资产凭证，用于代表艺术品、收藏品、游戏物品、虚拟地产等的所有权。
*   **稳定币**：价值与法币（如美元）锚定的加密货币，如USDC、DAI，是Web3世界重要的价值尺度和交换媒介。

### 4. 去中心化存储（Decentralized Storage）
区块链本身不适合存储大量数据（如图片、视频）。去中心化存储网络（如IPFS, Filecoin, Arweave）解决了这一问题。文件被分割、加密后分布式存储在全球多个节点上，用户通过内容寻址（Content Address）而非位置寻址（如URL）来获取数据，确保了数据的持久性和抗审查性。

### 5. 钱包（Wallets）
**加密货币钱包是用户进入Web3的入口和身份基石**。它不仅仅是一个存储资产的地方，更是一个**自我主权身份（Self-Sovereign Identity, SSI）管理器**。它主要管理两样东西：
*   **公钥**：相当于你的银行账号，可以公开分享，用于接收资产。
*   **私钥**：相当于你的银行卡密码+身份证明，必须绝对保密。谁拥有私钥，谁就真正拥有该地址下的所有资产和身份。

目前常见的钱包有MetaMask（浏览器插件）、Trust Wallet（手机端）等。

---

## 第三章：Web3生态全景与应用场景

Web3生态已经发展出一个庞大且多样化的应用体系，涵盖了金融、娱乐、文化、社会等多个领域。

### 1. DeFi（去中心化金融）
旨在构建一个向全球所有人开放、无需许可、透明透明的金融系统。核心应用包括：
*   **去中心化交易所**：如Uniswap, SushiSwap，允许用户直接点对点交易加密货币，无需中心化交易所撮合。
*   **借贷协议**：如Aave, Compound，用户可以将资产存入池中赚取利息，或通过超额抵押借出其他资产。
*   **稳定币**：如DAI，通过超额抵押加密资产生成与美元软锚定的稳定币。
*   **衍生品、保险、资产管理**等。

### 2. NFT与数字收藏品
NFT为数字世界的独一无二的事物提供了所有权证明。
*   **数字艺术**：艺术家可以直接向全球收藏家发售作品，并获得持续的版税收入。
*   **收藏品**：如 CryptoPunks, Bored Ape Yacht Club，形成了强大的社区和文化认同。
*   **游戏资产**：玩家真正拥有游戏内的道具、土地和角色，甚至可以跨游戏使用。
*   **身份与认证**：可用于代表会员资格、门票、学历证书等。

### 3. 元宇宙（Metaverse）与Web3游戏
Web3为元宇宙提供了理想的经济底层。虚拟世界（如Decentraland, The Sandbox）中的土地、物品均以NFT形式存在，由用户完全拥有。Play-to-Earn（边玩边赚）模式（如Axie Infinity）尝试将游戏产生的价值返还给玩家，重塑游戏经济模型。

### 4. DAO（去中心化自治组织）
DAO是一种基于区块链和智能合约的新型组织形态。没有传统的公司 hierarchy，规则通过代码制定，决策由社区通过治理代币投票完成。DAO可以被视为一个围绕共同目标运行的"互联网原生社区"。它被用于管理项目、投资基金、创建社区等。

### 5. 去中心化身份与社会（DID & DeSoc）
旨在解决Web2.0社交媒体的弊端。用户拥有一个可跨平台使用的、自主控制的数字身份（DID），其社交关系（Social Graph）也由自己掌握，不再被某个特定平台锁定。项目如Lens Protocol正在此领域进行探索。

---

## 第四章：挑战与批判性思考

Web3的发展并非一帆风顺，它面临着诸多严峻的挑战和质疑：

### 1. 可扩展性（Scalability）与用户体验
主流区块链（如以太坊）在处理大量交易时，常出现网络拥堵和交易费用（Gas Fee）高昂的问题。这严重影响了用户体验，阻碍大规模采用。Layer2扩容方案（如Optimism, Arbitrum）和分片技术正在努力解决此问题。

### 2. 安全风险
智能合约代码漏洞、私钥丢失或被盗、项目诈骗和 rug pull（拉地毯）等安全事件频发，给用户造成了巨大损失。安全审计和用户教育至关重要。

### 3. 监管不确定性
全球各国政府对加密货币和Web3的监管政策仍在探索和形成中。监管的不确定性可能抑制创新，但也可能为行业的健康发展提供必要的框架和消费者保护。

### 4. 中心化与去中心化的悖论
Web3生态中出现了新的中心化趋势，如中心化交易所（CEX）占据巨大流量，少数矿池/验证者掌握大量算力/质押代币。如何平衡效率与去中心化理想是一个长期课题。

### 5. 环境问题
基于工作量证明（PoW）共识的区块链（如早期的比特币、以太坊）能耗巨大，备受环保指责。向权益证明（PoS）等更节能共识机制的转型（如以太坊合并）正在缓解这一问题。

### 6. 投机性与价值泡沫
当前生态中存在大量投机行为，许多项目缺乏实际效用支撑，泡沫风险显著。真正的长期价值需要建立在解决实际问题和产生持续现金流的基础上。

---

## 第五章：未来展望与结语

Web3仍处于非常早期的阶段，相当于1990年代的互联网。它的最终形态远未定型，但其代表的方向——**赋予个体更多的权力、自由和经济能动性**——具有强大的吸引力。

未来的Web3发展可能会呈现以下趋势：

1.  **无缝的用户体验**：钱包、密钥管理等将变得更加易用和无感，降低普通用户的进入门槛。
2.  **与现实世界资产（RWA）的融合**：区块链将不仅用于纯数字资产，也将用于代表和交易房地产、股票、商品等现实世界资产。
3.  **合规化与制度化**：随着监管框架清晰，传统金融机构和大型科技公司将以更合规的方式进入该领域，带来巨大的资金和用户。
4.  **跨链互操作性**：不同的区块链网络将能够无缝通信和价值转移，形成一个真正的"价值互联网"。
5.  **探索非金融化应用**：Web3技术将在供应链管理、公共事务、医疗健康等更多领域找到应用场景。

**结语**

Web3是一场深刻的社会技术实验。它试图用密码学和分布式系统，来解决中心化机构带来的信任、公平和效率问题。它绝非仅仅是加密货币和投机，其内核是关于**数字时代的所有权、治理和人类协作方式**的重新构想。

这条路注定漫长且充满荆棘，技术瓶颈、人性贪婪和监管博弈将伴随其左右。但无论如何，Web3已经点燃了星星之火，挑战着我们对互联网的固有认知，并为我们描绘了一个更加开放、公平和赋予用户权力的数字未来。无论其最终成功与否，它都已经并将继续推动互联网技术和经济模型的创新边界。对于每一个互联网公民而言，理解Web3，就是试图理解我们可能即将步入的未来。
`
}
</script>

<template>
  <div v-if="article" class="article-content-page">
    <div class="article-content">
      <!-- 文章头部信息 -->
      <div class="article-header">
        <!-- <div class="article-cover">
          <img :src="article.cover" :alt="article.title" />
        </div> -->
        
        <h1 class="article-title">{{ article.title }}</h1>
        
        <div class="article-meta">
          <span class="meta-item date">{{ article.date }}</span>
          <span class="meta-item category">{{ article.category }}</span>
          <span class="meta-item views">{{ article.views }} 阅读</span>
        </div>
        
        <div class="article-tags">
          <span v-for="tag in article.tags" :key="tag" class="tag">{{ tag }}</span>
        </div>
      </div>
      
      <!-- 文章内容 -->
      <MarkdownRenderer :content="content" />
    </div>
    <!-- 底部粘性操作栏 -->
    <div class="sticky-action-bar">
      <button class="action-btn like-btn" @click="handleLike">
        <span class="icon">👍</span>
        <span class="text">点赞</span>
        <span class="count" v-if="likeCount > 0">{{ likeCount }}</span>
      </button>
      <button class="action-btn comment-btn" @click="scrollToComments">
        <span class="icon">💬</span>
        <span class="text">评论</span>
      </button>
    </div>
    
    <!-- 评论区组件 -->
    <Comments />
  </div>
    
  
  <div v-else class="article-not-found">
    <h2>文章不存在</h2>
    <p>抱歉，您请求的文章不存在或已被删除。</p>
    <router-link to="/articles" class="back-link">返回文章列表</router-link>
  </div>
</template>

<style scoped>
.article-content-page {
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  /* padding: 30px; */
  margin-bottom: 30px;
}

.article-content {
  margin: 30px 30px 0;
}

.article-header {
  text-align: left;
}

.article-cover {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
}

.article-cover img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.article-title {
  font-size: 32px;
  font-weight: 700;
  color: #2c3e50;
  margin-bottom: 16px;
  line-height: 1.3;
}

.article-meta {
  display: flex;
  gap: 16px;
  margin-bottom: 16px;
  color: #7f8c8d;
  font-size: 14px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 4px;
}

.article-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
  margin-bottom: 20px;
}

.tag {
  background-color: #f0f0f0;
  color: #555;
  padding: 4px 10px;
  border-radius: 4px;
  font-size: 12px;
}

/* GitHub Markdown 样式由 github-markdown-css 提供 */
.article-body {
  /* 保留必要的容器样式 */
  line-height: 1.8;
  font-size: 16px;
  text-align: left;
}

/* KaTeX 样式调整 */
.article-body .katex-display {
  margin: 1.5em 0;
  overflow-x: auto;
  overflow-y: hidden;
}

.article-not-found {
  text-align: center;
  padding: 50px 20px;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
}

.back-link {
  display: inline-block;
  margin-top: 20px;
  color: #3498db;
  text-decoration: none;
}

.back-link:hover {
  text-decoration: underline;
}

/* 底部粘性操作栏 */
.sticky-action-bar {
  position: sticky;
  bottom: 0px;
  display: flex;
  justify-content: flex-end;
  gap: 16px;
  margin-top: 40px;
  padding: 8px 12px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(10px);
  /* border-radius: 50px; */
  /* box-shadow: 0 1px 12px rgba(0, 0, 0, 0.15); */
  border: 1px solid rgba(255, 255, 255, 0.2);
  z-index: 20;
}

.action-btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border: none;
  border-radius: 25px;
  font-size: 14px;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.3s ease;
  background: #f8f9fa;
  color: #555;
}

.action-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.like-btn:hover {
  background: #e8f5e8;
  color: #27ae60;
}

.comment-btn:hover {
  background: #e8f4fd;
  color: #3498db;
}

.action-btn .icon {
  font-size: 16px;
}

.action-btn .count {
  background: rgba(0, 0, 0, 0.1);
  padding: 2px 6px;
  border-radius: 10px;
  font-size: 12px;
  min-width: 18px;
  text-align: center;
}



/* 响应式设计 */
@media (max-width: 768px) {
  .article-content-page {
    padding: 20px;
  }
  
  .article-cover {
    height: 200px;
  }
  
  .article-title {
    font-size: 24px;
  }

  .sticky-action-bar {
    bottom: 10px;
    padding: 10px 16px;
  }

  .action-btn {
    padding: 8px 16px;
    font-size: 13px;
  }

  .action-btn .icon {
    font-size: 14px;
  }
}

</style>