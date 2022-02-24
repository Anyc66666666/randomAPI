package sentence

import (
	"fmt"
	"math/rand"
)

func ST()string {

	answers := []string{
		"保持自己旁观者的清醒，永远不做局中人 ",
		"你平时珍惜的自行车骑都不舍得骑，别人却站起来蹬",
		"飞鸟与鱼不同路，从此山水不相逢",
		"无意即是有缘，偶然即是必然。不问客从何来，本站为缘而开",
		"高频率的聊天会让你产生恋爱的感觉",
		"你心里有答案了，不是吗",
		"悠悠岁月多漫长，生世轮回为谁茫？今日随缘来此处，尘封记忆在此敞。",
		"爱意随风起，风止意难平",
		"垂死病中惊坐起，忘了奖励我自己",
		"妙语连珠是猎物，支支吾吾是喜欢",
		"喜欢我望向别处时你落在我身上的目光",
		"大雾四起，我在无人处爱你",
		"当初先生退出文坛我是极力反对的",
		"她逼你看了吗？",
		"这周日你有空吗？",
		"你有没有觉得我们现在有点不一样了？",
		"分手吧",
		"我也不知道我还是不是爱你",
		"可能吧，不爱了吧",
		"我觉得你很幼稚",
		"我们互删吧",
		"你懂我意思了？",
		"我还喜欢你",
		"？",
		"哦",
		"嗯？",
		"就这样吧",
		"对不起...",
		"求求你了，别离开我",
		"我错了，我不应该的",
		"滚",
		"珍惜所有的不期而遇 看淡所有的不辞而别。",
		"我们都很擅长口是心非，却又希望对方有所察觉.",
		"那天车窗起雾 我在玻璃上写了你的名字 后来雾散了 你也没了",
		"根本没喜欢过我对吗？只是在你寂寞的时候，我刚好出现，你刚好来者不拒，一切都刚刚好啊，是这样吗？",
		"我没有想赢。只是不想输？",
		"你做饼干他品尝着，想起这画面的揪心我隐藏着",
		"这世上真话本就不多，一个女子的脸红胜过一大段对白，可后来有了胭脂，便分不清是真情还是假意。———老舍《骆驼祥子》",
		" 当你停止主动的时候，这段关系就结束了 ",
		" 每一个决定转身的人 都在风里站了很久. ",
		"当初，我挽留你时说的话，连我自己都看不起自己了",
		"你在我身边的时候你就是一切 你不在我身边的时候一切都是你",
		"我没拒绝吹来的风，也再没有喜欢的人。这次我站在雾里，连自己都看不清",
		"我坐在湖边 以为是在海边 我吹着湖边的风也以为是海风 我喜欢你也以为你喜欢我 但湖是湖 海是海",
		"﹝她满脸嫌弃的求我放过她，我忽然愣住了。想起了很久以前小心翼翼的问我， 你会离开我吗！ ﹞",
		"而你心房的新房客，陪你欣赏夕阳的金黄色。",
		"心决定的事情，嘴巴又怎么讲的清楚。",
		" 天下熙熙，皆为利来;天下攘攘，皆为利往",
		"曹操乃是，一代枭雄。曹阿满。是一个无匹流氓。曹贼乃是一种精神。而我却占了两样曹阿满和曹贼的一种精神",
		"我真搞不懂他们这些人，年纪轻轻的非要去干这个，道德在哪里，底线在哪里，良知在哪里！地址在哪里！地址在哪里！",
		"生抽提鲜，老抽入味",
		"你们的幸福可以小点声吗？吵到我一个人的孤独了",
		" 不是所有人都会出轨，也不是所有人都会老老实实，但总会有会为你守住底线的人，不要被现在的风气蒙蔽了眼睛，爱你的人不会背叛你。",
		"不得不说，这些女孩是真的漂亮。气质出众魅力十足，尤其是那眼睛，不多不少，正好两个。",
		"以色交者 色衰则疏 以貌交者 久之则逆。唯有以心交者方能永恒",
		"我说真的，其实咱们两都知道这不叫喜欢也不叫爱，只不过是因为咱们两最近聊天太频繁了，让你感觉到那种喜欢或者说是新鲜感，其实真真的感情是抑制不了的，就算你不说话，你的眼睛你的心跳也会出卖你，喜欢一个人真的忍不住，可能我们都太着急了，打了几次电话，说了几次晚安，但是我认为爱情还是需要长久的沉淀吧，轰动并不是勇敢的爱，我希望我们都可以冷静下来好好想想，我还是希望我们多去了解了解吧，我说这么多也不是拒绝你什么，我只是希望对你负责当然了也是对我自己负责",
		"我给你点个赞，希望你能主动联系我，你主动我们就会有故事，我这是在通知你而不是在求你，希望你不要不知好歹",
		"有没有一种可能，她得了癌症不想拖累你？",
		"枕头朝哪边才能做这个梦",
		"我从未拥有过她，我只是陪她聊了很久很久",
		"有的人不冷，但偏偏喜欢烤火的感觉，便使劲往里边加柴，火更大了，人一热，就坐的远远的，本来火苗足以燃烧很久的柴被那人浪费了，没多久熄灭了",
		" 其实有些事情只有你一个人在遗憾 ",
		" 我原本以为吕布已经天下无敌了，没想到有人比他还勇猛，这是谁的部将？",
		"你本来就无人问津，何来的销声匿迹。",
		" 当你停止主动的时候，这段关系就结束了 ",
		"人生的路，靠自己一步步走去，真正能保护你的，是你自己的人格选择和文化选择。那么反过来，真正能伤害你的，也是一样，自己的选择。",
		" 非礼勿视，非礼勿听，非礼勿言，非礼勿动",
		"风吹包谷杆，杆杆碰杆杆",
		" 村口的狗叫了  其它的狗也叫了  但是它们不知道为什么叫",
		"人是不能闲的，一闲就会感情泛滥。所谓矫情屁话多，空虚寂寞冷，都是因为懒散堕落闲。",
		"这么漂亮的小姐姐一想不是我的，气的我一巴掌摔在我的法拉利上，然后掉出来三节电池",
		"树是生活，埋的是我，看花就好，别看我落魄！",
		"两处相思同沐雪，此生也算共白头",
		" 所幸思念无声从未惊扰旁人. ",
		" 当你停止主动的时候，这段关系就结束了 ",
		" 路灯下，地上有两个影子，这个是我的，另一个也是我的 ",
		"挡风玻璃为什么比后视镜大，因为前面的路更重要",
		" 分手后 对方单身的时间 就是对这段感情的评价 ",
		" 你平时珍惜的自行车骑都不舍得骑，别人却站起来蹬 ",
		"大抵是到了该寻一个姑娘的年纪了，近来夜里冷的厉害，特别是心里，凉的出奇，两床被子面对这寒冬的挑衅，也显得有些许吃力了，或许只有心仪姑娘的照料，才能让我感到温暖了罢",
		"谈个恋爱还想要亲嘴？那结了婚还不得在一起睡？",
		" 秋不秋天我不知道反正该黄的都黄了，搓搓手又要一个人过冬天了",
		"对，我可以惯着你，也可以换了你",
		"错过了夏天的爱情，没喝到秋天的奶茶，但我始终相信，冬天的西北风一定有我一口",
		" 你和我在一起，是不是因为他喜欢人妻？ ",
		"其实真正的有钱人是很低调的，外表是看不出来的。就拿我来说，虽然我在外面经常骑小黄车，但谁又会知道其实晚上睡觉电风扇都是开三挡。",
		"孤独是什么！孤独二字拆开来看，有孩童，有瓜果，有小犬，有蝴蝶，稚儿擎瓜柳打掤下，细犬逐蝶窄巷中，人间繁华多笑语，唯我空余两鬓风，孩童水果猫狗飞蝶当然热闹，可都与我无关，这就叫孤独",
		"倘是往日，倒也能呼头大睡，可今夜却如坐针毡，沉思良久，许是忘了奖励自己了罢",
		"为什么不找我聊天，我专门为你学的打字",
		" 学生时代的遗憾大概是：她身上的校服总是香的 她说那是洗衣液的味道 后来我用遍了所有洗衣液 都也没闻到这个味道",
		"卖过草鞋吃过苦，打过曹魏抗过吴。 终是皇叔负了蜀，接着奏乐接着舞。",
		"世间女子本就是一朵花，没有丑与美之分，只要落到心上人眼里，便是无暇之花，惊鸿一瞥",
		"我可以和一千个人做朋友,可以对一百个人微笑,对十个人讲情话,但是我的心只能给你一个人,听得到我对你爱的呼唤了吗",
		"热恋期的话都是屁话，如果在一起很久了当我是屁话",
		"女孩子太晚了就不要穿的花枝招展的出去乱逛。 会被变态尾随的。很不安全。我昨晚就尾随了一个",
		"我已经闭好一只眼睛 等你和我说晚安 我就闭上另一只",
		"抛硬币吧，抛第二次的时候你就有答案了",
		"大概是我嬉笑太甚，所以无人欣赏我的认真",
		"狂热直白的爱不丢脸，玩弄别人的感情才应该感到耻辱",
		"假如鹦鹉在古代灭绝，史书记载：一鸟羽毛七彩，会人语，你相信吗？",
		"瞻顾遗迹，如在昨日，令人长号不自禁",
		"其实好多时候，故事开始，你跟我都是一个人，但也有可能因为一次遇见变成我们俩",
		"男人遇到什么样的女人都不要怂，拼个未来给她，给不了提升了自己也是好事",
		"相见不如怀念，见了可能就不是以前的模样了，那份小美好，还是留在心里",
		"夏虫不可语冰",
		"其实好多时候，你没有办法得到一个东西的时候，你唯一可以做的就是，命令自己不可以忘记",
		"二十几岁的你，迷茫又着急。你想要房子你想要汽车，你想要旅行你想要高品质生活。你那么年轻却窥觑整个世界，你那么浮躁却想要看透生活。你不断催促自己赶快成长，却沉不下心来认真读一篇文章。你一次次吹响前进的号角，却总是倒在离出发不远的地方。请坚韧一些，年轻的朋友，勇敢地前行。",
		"很多年了 只还记得她的名字  已经忘了她的样子",
		"被爱者风生水起，追爱者一事无成！",
		"一个女孩子可以陪你吃地摊骑小电驴但你不能认为她只配吃地摊骑小电驴 共勉兄弟们",
		"地球是圆的，你送出去的每一颗糖，最后都会回到你手上",
		"我只是你的故人  却不是你故事里的人",
		"明明自己过得也不尽如人意，可偏偏见不得这人间疾苦",
		"文案撩的狠 追你又不肯  私聊让我滚 无情又残忍",
		"孤独本是常态，逢人不必深言.",
		" 我幼稚的时候好爱你啊 而你大概只看到了幼稚",
		"余则单身成狗处其间，略无慕艳意",
		" 别让无趣的世俗淹没了你的浪漫和热情 ",
		"今天早上在厨房发现了一只蚂蚁 我在它面前放了块糖 它碰了碰，跑回家叫同伴了 我迅速的藏起了这块糖 因为我想让它朋友都觉得它是骗子",
		" 人总在睡前想一些念念不忘的事. ",
		"本来挺喜欢你的，但是突然看到了不该看见的东西，知道了一些不该知道的事情，我就突然觉得你和别人好像都一样。就突然没意思了",
		"有一个人教会了你什么是爱，但是，她却不爱你了",
	}
	fmt.Println(answers[rand.Intn(len(answers))])
	return answers[rand.Intn(len(answers))]
}