var cipai = '入塞河传三台解红塞姑韵令侧犯尾犯天香步月暗香大有催雪国香大椿花犯眉妩索酒南浦西河秋霁薄幸疏影宣清八归六丑多丽戚氏'
LEN = 2
let buffer = ''
let c: string = ''

for (let i = 0; i < cipai.length; i++) {
	if (i % LEN == 0) {
		console.log(buffer)
		buffer = cipai[i]
	} else {
		buffer += cipai[i]
	}

}