#include <stdio.h>
#include <stdlib.h>
#define MaxRuleNum 10
#define MaxVnNum 10
#define MaxVtNum 10
#define MaxStackDepth 20
#define MaxPLength 20
#define MaxStLength 50
struct pRNode { //产生式右部的结构
	int rCursor;//指向字符表的指针表示产生式右部字符
	struct pRNode *next;//链表的组织结构
};
struct pNode {
	int lCursor; //指向产生式左边字符
	int rLength; //产生式右部分的长度
	struct pRNode *rHead; //产生式右部分头指针
};

char Vn[MaxVnNum + 1]; //非终结符集
int vnNum;

char Vt[MaxVtNum + 1]; // 终结符集 
int vtNum;

struct pNode P[MaxRuleNum];//产生式集
int PNum;

char buffer[MaxPLength + 1];//缓冲串集合
char ch;

char st[MaxStLength]; //要分析的符号串

struct collectNode { //定义first和follow集
	int nVt;//字符
	struct collectNode *next;
};

struct collectNode* first[MaxVnNum + 1]; //first 集每个元素是个指向链表的指针
struct collectNode* follow[MaxVnNum + 1]; //follow 集

int analyseTable[MaxVnNum + 1][MaxVtNum + 1 + 1];//分析表
int analyseStack[MaxStackDepth + 1]; //分析栈
int topAnalyse; /* 分析栈顶 */
void Init();/* 初始化 */
int IndexCh(char ch);
void InputVt(); /* 输入终结符 */
void InputVn();/* 输入非终结符 */
void ShowChArray(char* collect, int num);/* 输出 Vn 或 Vt 的内容 */
void InputP();/* 产生式输入 */
bool CheckP(char * st);/* 判断产生式正确性 */
void First(int U);
void AddFirst(int U, int nCh); /* 加入 first 集*/
bool HaveEmpty(int nVn);
void Follow(int V);/* 计算 follow 集*/
void AddFollow(int V, int nCh, int kind);
void ShowCollect(struct collectNode **collect);/* 输出 first 或 follow集*/
void FirstFollow();/* 计算 first 和 follow*/
void CreateAT();/* 构造预测分析表 */
void ShowAT();/* 输出分析表 */
void Identify(char *st);
void InitStack();
void ShowStack();
void Pop();
void Push(int r);
void printp();//输出所有产生式
void isdturn();//消除直接左递归
void isjturn();//消除间接左递归
void length();//更新长度

int main() {
	char todo, ch;
	Init();
	InputVn();
	InputVt();
	InputP();
	getchar();//吃掉换行
	isjturn();
	isdturn();
	length();
	printf(" 所得不含左递归的产生式为：\n");
	printp();
	FirstFollow();
	printf(" 所得 first 集为： ");
	ShowCollect(first);
	printf(" 所得 follow 集为： ");
	ShowCollect(follow);
	CreateAT();
	ShowAT();
	todo = 'y';
	while ('y' == todo) {
		printf("\n 是否继续进行句型分析？ (y / n):");
		todo = getchar();
		while ('y' != todo && 'n' != todo) {
			printf("\n(y / n)? ");
			todo = getchar();
		}
		if ('y' == todo) {
			int i;
			InitStack();
			printf(" 请输入符号串 ( 以#结束 ) : ");
			ch = getchar();
			i = 0;
			while ('#' != ch && i < MaxStLength) {
				if (' ' != ch && '\n' != ch) {
					st[i++] = ch;
				}
				ch = getchar();
			}
			if ('#' == ch && i < MaxStLength) {
				st[i] = ch;
				Identify(st);
			}
			else
				printf(" 输入出错！ \n");
		}
	}
	getchar();
}
void Init() {
	int i, j;
	vnNum = 0;
	vtNum = 0;
	PNum = 0;
	for (i = 0; i <= MaxVnNum; i++)
		Vn[i] = '\0';//非终初始化
	for (i = 0; i <= MaxVtNum; i++)
		Vt[i] = '\0';//终初始化
	for (i = 0; i < MaxRuleNum; i++) {//产生式初始化
		P[i].lCursor = NULL;
		P[i].rHead = NULL;
		P[i].rLength = 0;
	}
	for (i = 0; i <= MaxPLength; i++)
		buffer[i] = '\0';//字符集初始化
	for (i = 0; i < MaxVnNum; i++) {
		first[i] = NULL;//first初始化
		follow[i] = NULL;//follow初始化
	}
	for (i = 0; i <= MaxVnNum; i++) {
		for (j = 0; j <= MaxVnNum + 1; j++)
			analyseTable[i][j] = -1;//分析表的初始化
	}
}
int IndexCh(char ch) {
	int n;
	n = 0; /*is Vn?*/
	while (ch != Vn[n] && '\0' != Vn[n])
		n++;
	if ('\0' != Vn[n])
		return 100 + n;//非终的值是大于100的
	n = 0; /*is Vt?*/
	while (ch != Vt[n] && '\0' != Vt[n])
		n++;
	if ('\0' != Vt[n])
		return n;
	return -1;//空返回-1
}

void ShowChArray(char* collect) {
	int k = 0;
	while ('\0' != collect[k]) {
		printf(" %c ", collect[k++]);
	}
	printf("\n");
}

void InputVn() {
	int inErr = 1;
	int n, k;
	char ch;
	while (inErr) {
		printf(" 请输入所有的非终结符，注意： ");
		printf(" 请将开始符放在第一位，并以 #号结束 :\n");
		ch = ' ';
		n = 0;
		/* 初始化数组 */
		while (n < MaxVnNum) {
			Vn[n++] = '\0';
		}
		n = 0;
		while (('#' != ch) && (n < MaxVnNum)) {
			if (' ' != ch && '\n' != ch && -1 == IndexCh(ch)) {//排除空格,换行,已经输入的字符
				Vn[n++] = ch;
				vnNum++;
			}
			ch = getchar();
		}
		Vn[n] = '#'; /* 以"#" 标志结束用于判断长度是否合法 */
		k = n;
		if ('#' != ch) {
			if ('#' != (ch = getchar())) {
				while ('#' != (ch = getchar()))
					;//清除缓冲区,一会重新输入不能在读进去了
				printf("\n 符号数目超过限制！ \n");
				inErr = 1;
				continue;
			}
		}
		/* 正确性确认，正确则，执行下下面，否则重新输入 */
		Vn[k] = '\0';
		ShowChArray(Vn);
		ch = ' ';
		while ('y' != ch && 'n' != ch) {
			if ('\n' != ch) {
				printf(" 输入正确确认 ?(y/n):");
			}
			scanf("%c", &ch);
		}
		if ('n' == ch) {
			printf(" 录入错误重新输入！ \n");
			inErr = 1;
		}
		else {
			inErr = 0;
		}
	}
}
/* 输入终结符 */
void InputVt() {
	int inErr = 1;
	int n, k;
	char ch;
	while (inErr) {
		printf("\n 请输入所有的终结符，注意： ");
		printf(" 以#号结束 :\n");
		ch = ' ';
		n = 0;
		/* 初始化数组 */
		while (n < MaxVtNum) {
			Vt[n++] = '\0';
		}
		n = 0;
		while (('#' != ch) && (n < MaxVtNum)) {
			if (' ' != ch && '\n' != ch && -1 == IndexCh(ch)) {
				Vt[n++] = ch;
				vtNum++;
			}
			ch = getchar();
		}
		Vt[n] = '#';//默认最后一个是#
		k = n;
		if ('#' != ch) {
			if ('#' != (ch = getchar())) {
				while ('#' != (ch = getchar()))
					;
				printf("\n 符号数目超过限制！ \n");
				inErr = 1;
				continue;
			}
		}
		Vt[k] = '\0';
		ShowChArray(Vt);
		ch = ' ';
		while ('y' != ch && 'n' != ch) {
			if ('\n' != ch) {
				printf(" 输入正确确认 ?(y/n):");
			}
			scanf("%c", &ch);
		}
		if ('n' == ch) {
			printf(" 录入错误重新输入！ \n");
			inErr = 1;
		}
		else {
			inErr = 0;
		}
	}
}
/* 产生式输入 */
void InputP() {
	char ch;
	int i = 0, n, num;
	printf(" 请输入文法产生式的个数： ");
	scanf("%d", &num);
	PNum = num;
	getchar(); //删除回车
	printf("\n 请输入文法的 %d个产生式 , 并以回车分隔每个产生式： ", num);
	printf("\n");
	while (i < num) {
		printf(" 第%d 个： ", i);
		/* 初始化 */
		for (n = 0; n < MaxPLength; n++)
			buffer[n] = '\0';
		/* 输入产生式串 */
		ch = ' ';
		n = 0;
		while ('\n' != (ch = getchar()) && n < MaxPLength) {
			if (' ' != ch)//除去空格换行
				buffer[n++] = ch;
		}
		buffer[n] = '\0';
		if (CheckP(buffer)) {
			pRNode *pt, *qt;
			P[i].lCursor = IndexCh(buffer[0]);//产生式左边赋值
			pt = (pRNode*)malloc(sizeof(pRNode));//构造产生式右部
			pt->rCursor = IndexCh(buffer[3]);
			pt->next = NULL;
			P[i].rHead = pt;//构造头结点
			n = 4;
			while ('\0' != buffer[n]) {//构造链表
				qt = (pRNode*)malloc(sizeof(pRNode));
				qt->rCursor = IndexCh(buffer[n]);
				qt->next = NULL;
				pt->next = qt;
				pt = qt;//pt跟着qt
				n++;
			}
			P[i].rLength = n - 3;//产生式右部长度
			i++;
		}
		else
			printf(" 输入符号含非法在成分，请重新输入！ \n");
	}
}
/* 判断产生式正确性 */
bool CheckP(char * st) {
	int n;
	if (100 > IndexCh(st[0]))//如果第一个字符是终结符
		return false;
	if ('-' != st[1])//终结符后面不是->则报错
		return false;
	if ('>' != st[2])
		return false;
	for (n = 3; '\0' != st[n]; n++) {
		if (-1 == IndexCh(st[n]))//输入了不存在的报错
			return false;
	}
	return true;
}
//消除间接左递归
void isjturn()
{
	for (int j = 0; j<vnNum; j++)
	{
		for (int i = 0; i < PNum; i++)
		{
			if (j == (P[i].lCursor - 100) && P[i].rHead->rCursor >= 100 && P[i].rHead->rCursor<P[i].lCursor)
			{
				for (int k = 0; k<i; k++)
				{
					if (P[k].lCursor == P[i].rHead->rCursor)
					{
						pRNode * p = P[k].rHead;
						pRNode * qt = (pRNode*)malloc(sizeof(pRNode));
						pRNode * head = qt;
						while (p != NULL)
						{

							pRNode * pt = (pRNode*)malloc(sizeof(pRNode));
							pt->rCursor = p->rCursor;
							qt->next = pt;
							qt = pt;
							p = p->next;
						}
						qt->next = P[i].rHead->next;
						P[i].rHead = head->next;
					}

				}

			}
		}
	}
}



//消除直接左递归
void isdturn()
{

	int flag2 = 0;

	struct collectNode *pa[10];//作为a的数组
	struct collectNode *pb[10];
	for (int i = 0; i<10; i++)
	{
		pa[i] = pb[i] = NULL;
	}
	int n = 0, m = 0;
	pRNode * p = NULL;
	for (int j = 0; j<vnNum; j++)
	{
		int flag1 = 0;
		for (int i = 0; i < PNum; i++)
		{

			if (j == P[i].lCursor - 100 && j == P[i].rHead->rCursor - 100)
			{
				flag1 = 1; break;
			}
		}
		if (flag1 == 1)
		{
			for (int i = 0; i < PNum; i++)
			{
				if (j == (P[i].lCursor - 100) && j == (P[i].rHead->rCursor - 100))
				{
					n++;
					P[i].lCursor = vnNum + 100;
					p = P[i].rHead = P[i].rHead->next;
					pRNode* pt = (pRNode*)malloc(sizeof(pRNode));//构造产生式右部
					pt->rCursor = vnNum + 100;
					pt->next = NULL;
					while (p->next != NULL) p = p->next;
					p->next = pt;
				}
				if (j == (P[i].lCursor - 100) && j != (P[i].rHead->rCursor - 100) && -1 != P[i].rHead->rCursor)
				{
					p = P[i].rHead = P[i].rHead;
					pRNode* pt = (pRNode*)malloc(sizeof(pRNode));//构造产生式右部
					pt->rCursor = vnNum + 100;
					pt->next = NULL;
					while (p->next != NULL) p = p->next;
					p->next = pt;
				}
			}
			P[PNum++].lCursor = vnNum + 100;
			pRNode* pt = (pRNode*)malloc(sizeof(pRNode));//构造产生式右部
			pt->rCursor = -1; pt->next = NULL;
			P[PNum - 1].rHead = pt;
			Vn[vnNum++] = Vn[j] + 32;
		}

	}
}

void First(int U) {
	int i, j;
	for (i = 0; i < PNum; i++) {//对于每一个产生式
		if (P[i].lCursor == U) {//如果产生式的左边是该非
			struct pRNode* pt;
			pt = P[i].rHead;//找到右部产生式
			j = 0;
			while (j < P[i].rLength) {
				if (100 > pt->rCursor) {//如果是终结符
					AddFirst(U, pt->rCursor);//加到它的first集合就不在往后看了
					break;
				}
				else {
					if (NULL == first[pt->rCursor - 100]) {//非终结符如果它first集合空
						First(pt->rCursor);//求他的first集合
					}
					AddFirst(U, pt->rCursor);//把它的first集合加入,非空
					if (!HaveEmpty(pt->rCursor)) {
						break;
					}
					else {
						pt = pt->next;//如果该非终的first有空找下一个
					}
				}
				j++;
			}
			if (j >= P[i].rLength) // 当产生式右部都能推出空时 才把空加进去
				AddFirst(U, -1);
		}
	}
}

/* 加入 first 集*/
void AddFirst(int U, int nCh) {//把nch加到终结符U的first集
	struct collectNode *pt, *qt;
	int ch; /* 用于处理 Vn*/
	pt = NULL;
	qt = NULL;
	if (nCh < 100) {//对于终结符
		pt = first[U - 100];//找到它的first集合的链表
		while (NULL != pt) {
			if (pt->nVt == nCh)
				break;
			else {
				qt = pt;
				pt = pt->next;//qt记录前一,pt一直到空
			}
		}
		if (NULL == pt) {
			pt = (struct collectNode *)malloc(sizeof(struct collectNode));
			pt->nVt = nCh;
			pt->next = NULL;
			if (NULL == first[U - 100]) {//如果本来就没有,则直接将他作为头结点
				first[U - 100] = pt;
			}
			else {
				qt->next = pt; //新加入到first集合的字符连到链表上
			}
			pt = pt->next;
		}
	}
	else {//终结符
		pt = first[nCh - 100];
		while (NULL != pt) {
			ch = pt->nVt;
			if (-1 != ch) {
				AddFirst(U, ch);//将它first集合非空部分加入。
			}
			pt = pt->next;
		}
	}
}
bool HaveEmpty(int nVn) {//非终且能推出空则true
	if (nVn < 100)
		return false;
	struct collectNode *pt;
	pt = first[nVn - 100];
	while (NULL != pt) {
		if (-1 == pt->nVt)
			return true;
		pt = pt->next;
	}
	return false;
}
void Follow(int V) {
	int i;
	struct pRNode *pt;
	if (100 == V) // 当为字符是初始字符的时候
		AddFollow(V, -1, 0);//直接把结束符加进去
	for (i = 0; i < PNum; i++) {
		pt = P[i].rHead;
		while (NULL != pt && pt->rCursor != V)
			pt = pt->next;//产生式的有部找V
		if (NULL != pt) {//找到了
			pt = pt->next;
			if (NULL == pt) {//是最后一个字符
				if (NULL == follow[P[i].lCursor - 100] && P[i].lCursor != V) {//把左侧非V且有follow的非的集合加进去(含空)
					Follow(P[i].lCursor);
				}
				AddFollow(V, P[i].lCursor, 0);
			}
			else {
				while (NULL != pt && HaveEmpty(pt->rCursor)) {//因为没有判断fellow集合是否不变的机制,所以需要一直看到头
					AddFollow(V, pt->rCursor, 1);//如果后面是非终有空
					pt = pt->next;//每一个非终的first
				}
				if (NULL == pt) {//如果后面没有了
					if (NULL == follow[P[i].lCursor - 100] && P[i].lCursor != V) {
						Follow(P[i].lCursor);//那么说明这个非终是最后一个,那么把产生式左边的fello加进去
					}
					AddFollow(V, P[i].lCursor, 0);
				}
				else {
					AddFollow(V, pt->rCursor, 1);//否则最后一个姚明是终要么是不含空非终first加入
				}
			}
		}
	}
}
void AddFollow(int V, int nCh, int kind) {
	struct collectNode *pt, *qt;
	int ch;
	pt = NULL;
	qt = NULL;
	if (nCh < 100) { /* 为终结符时直接加入链表尾部 */
		pt = follow[V - 100];
		while (NULL != pt) {
			if (pt->nVt == nCh)
				break;
			else {
				qt = pt;
				pt = pt->next;
			}
		}
		if (NULL == pt) {
			pt = (struct collectNode *)malloc(sizeof(struct collectNode));
			pt->nVt = nCh;
			pt->next = NULL;
			if (NULL == follow[V - 100]) {
				follow[V - 100] = pt;
			}
			else {
				qt->next = pt;
			}
			pt = pt->next;
		}
	}
	else {//非终结符的时候
		if (0 == kind) {//直接把follow加进去
			pt = follow[nCh - 100];
			while (NULL != pt) {
				ch = pt->nVt;
				AddFollow(V, ch, 0);
				pt = pt->next;
			}
		}
		else {//把不含空的first加进去
			pt = first[nCh - 100];
			while (NULL != pt) {
				ch = pt->nVt;
				if (-1 != ch) {
					AddFollow(V, ch, 1);
				}
				pt = pt->next;
			}
		}
	}
}
/* 输出 first 或 follow 集*/
void ShowCollect(struct collectNode **collect) {
	int i;
	struct collectNode *pt;
	i = 0;
	while (NULL != collect[i]) {
		pt = collect[i];
		printf("\n%c:\t", Vn[i]);
		while (NULL != pt) {
			if (-1 != pt->nVt) {
				printf(" %c", Vt[pt->nVt]);
			}
			else
				printf(" #");//把空转化成"#"
			pt = pt->next;
		}
		i++;
	}
	printf("\n");
}
/* 计算 first 和 follow*/
void FirstFollow() {
	int i;
	i = 0;
	while ('\0' != Vn[i]) {
		if (NULL == first[i])
			First(100 + i);//非终结符对应的序号在100外
		i++;
	}
	i = 0;
	while ('\0' != Vn[i]) {
		if (NULL == follow[i])
			Follow(100 + i);
		i++;
	}
}
/* 构造预测分析表 */
void CreateAT() {
	int i;
	struct pRNode *pt;
	struct collectNode *ct;
	for (i = 0; i < PNum; i++) {//对每个产生式
		pt = P[i].rHead;
		while (NULL != pt && HaveEmpty(pt->rCursor)) {//非终且能推出空
			ct = first[pt->rCursor - 100];
			while (NULL != ct) {
				if (-1 != ct->nVt)
					analyseTable[P[i].lCursor - 100][ct->nVt] = i;//产生式左侧遇到它的右侧第一个非终的first集合
				ct = ct->next;
			}
			pt = pt->next;
		}
		if (NULL == pt) {//后面都可以推出空
			ct = follow[P[i].lCursor - 100];
			while (NULL != ct) {
				if (-1 != ct->nVt)
					analyseTable[P[i].lCursor - 100][ct->nVt] = i;
				else//follow空其实就是遇到了#
					analyseTable[P[i].lCursor - 100][vtNum] = i;//最后一列是#
				ct = ct->next;
			}
		}
		else {//最后一个不能推出空
			if (100 <= pt->rCursor) { /* 不含空的非终结符 */
				ct = first[pt->rCursor - 100];
				while (NULL != ct) {
					analyseTable[P[i].lCursor - 100][ct->nVt] = i;
					ct = ct->next;
				}
			}
			else { /* 非空的终结符或者空 */
				if (-1 == pt->rCursor) {//相当于是推出了空
					ct = follow[P[i].lCursor - 100];
					while (NULL != ct) {
						if (-1 != ct->nVt)
							//if(analyseTable[P[i].lCursor - 100][ct->nVt] ==-1 )
							analyseTable[P[i].lCursor - 100][ct->nVt] = i;
						else /* 当含有 # 号时 */
							analyseTable[P[i].lCursor - 100][vtNum] = i;
						ct = ct->next;
					}
				}
				else { /* 为终结符 */
					analyseTable[P[i].lCursor - 100][pt->rCursor] = i;
				}
			}
		}
	}
}
/* 输出分析表 */
void ShowAT() {

	int i, j;
	printf(" 构造预测分析表如下： \n");
	printf("\t|\t");
	for (i = 0; i < vtNum; i++) {
		printf("%c\t", Vt[i]);//行头是所有终
	}
	printf("#\t\n");//#单独输出
	printf("- - -\t|- - -\t");
	for (i = 0; i <= vtNum; i++)
		printf("- - -\t");
	printf("\n");
	for (i = 0; i < vnNum; i++) {
		printf("%c\t|\t", Vn[i]);//列头是终结符
		for (j = 0; j <= vtNum; j++) {
			if (-1 != analyseTable[i][j])
				printf("R(%d)\t", analyseTable[i][j]);
			else
				printf("     \t");
		}
		printf("\n");
	}
}

void Identify(char *st) {//构造栈
	int current, step, r; /*r 表使用的产生式的序号 */
	printf("\n%s 的分析过程： \n", st);
	printf(" 步骤 \t 分析符号栈 \t 当前指示字符 \t 使用产生式序号 \n");
	step = 0;//第几步
	current = 0;//输入串指针
	printf("%d\t", step);
	ShowStack();//栈你初始就有#和开始符号
	printf("\t\t%c\t\t- -\n", st[current]);
	while ('#' != st[current]) {//当输入串剩下#结束
		if (100 > analyseStack[topAnalyse]) {
			if (analyseStack[topAnalyse] == IndexCh(st[current])) {
				Pop();//终结符匹配,则出栈
				current++;
				step++;
				printf("%d\t", step);
				ShowStack();
				printf("\t\t%c\t\t 出栈、后移 \n", st[current]);
			}
			else {
				printf("%c-%c 不匹配！ ", analyseStack[topAnalyse], st[current]);
				printf(" 此串不是此文法的句子！ \n");
				return;
			}
		}
		else { /* 当为非终结符时 */
			r = analyseTable[analyseStack[topAnalyse] -
				100][IndexCh(st[current])];
			if (-1 != r) {
				Push(r);//push里有pop
				step++;
				printf("%d\t", step);
				ShowStack();
				printf("\t\t%c\t\t%d\n", st[current], r);
			}
			else {
				printf(" 此串不是此文法的句子！ \n");
				return;
			}
		}

	}
	if ('#' == st[current]) {
		if (0 == topAnalyse && '#' == st[current]) {
			step++;
			printf("%d\t", step);
			ShowStack();
			printf("\t\t%c\t\t 分析成功！ \n", st[current]);
			printf("%s 是给定文法的句子！ \n", st);
		}
		else {
			while (topAnalyse > 0) {
				if (100 > analyseStack[topAnalyse]) {//终结符肯定不是
					printf(" 此串不是此文法的句子！ \n");
					return;
				}
				else {//虽然输入串此时是#当站内是非终是,只有遇到推出空才行
					r = analyseTable[analyseStack[topAnalyse] - 100][vtNum];
					if (-1 != r) {
						Push(r);
						step++;
						printf("%d\t", step);
						ShowStack();
						if (0 == topAnalyse && '#' == st[current]) {//可能有多个符号条件的非
							printf("\t\t%c\t\t 分析成功！ \n", st[current]);
							printf("%s 是给定文法的句子！ \n", st);
						}
						else
							printf("\t\t%c\t\t%d\n", st[current], r);
					}
					else {
						printf(" 此串不是此文法的句子！ \n");
						return;
					}
				}
			}
		}
	}
}
/* 初始化栈及符号串 */
void InitStack() {
	int i;
	/* 分析栈的初始化 */
	for (i = 0; i < MaxStLength; i++)
		st[i] = '\0';
	analyseStack[0] = -1; /*#(-1) 入栈 */
	analyseStack[1] = 100; /* 初始符入栈 */
	topAnalyse = 1;//栈顶指针,指向栈顶元素
}
/* 显示符号栈中内容 */
void ShowStack() {
	int i;
	for (i = 0; i <= topAnalyse; i++) {
		if (100 <= analyseStack[i])
			printf("%c", Vn[analyseStack[i] - 100]);
		else {
			if (-1 != analyseStack[i])
				printf("%c", Vt[analyseStack[i]]);
			else
				printf("#");
		}
	}
}
/* 栈顶出栈 */
void Pop() {
	topAnalyse--;
}
void Push(int r) {
	int i;
	struct pRNode *pt;
	Pop();//
	pt = P[r].rHead;
	if (-1 == pt->rCursor)
		return;//产生式右部是空不进栈
	topAnalyse += P[r].rLength;
	for (i = 0; i < P[r].rLength; i++) {
		analyseStack[topAnalyse - i] = pt->rCursor;/* 逆序入栈 */
		pt = pt->next;
	}
}
//输出所有产生式
void printp()
{
	for (int i = 0; i < PNum; i++)
	{
		putchar(Vn[P[i].lCursor - 100]);
		printf("->");
		pRNode* pt = P[i].rHead;
		while (pt != NULL)
		{
			if (pt->rCursor >= 100)
				putchar(Vn[pt->rCursor - 100]);
			else putchar(Vt[pt->rCursor]);
			pt = pt->next;
		}
		printf("\n");
	}
}

void length()//更新有部产生式的长度
{
	for (int i = 0; i < PNum; i++)
	{
		int count = 0;
		pRNode* pt = P[i].rHead;
		while (pt != NULL)
		{
			pt = pt->next;
			count++;
		}
		P[i].rLength = count;
	}

}


