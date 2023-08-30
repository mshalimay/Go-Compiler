source_filename = "binarytree"
target triple = "x86_64-linux-gnu"

%struct.Node = type {i64, %struct.Node*, %struct.Node*}
@.nilNode = common global %struct.Node* null

@root = common global %struct.Node* null

define i64 @compare(i64 %cur, i64 %neuw)
{
L0:
	%compare_retval = alloca i64
	%P_cur = alloca i64
	; parameters
	%P_neuw = alloca i64
	; local variables
	store i64 %cur, i64* %P_cur
	store i64 %neuw, i64* %P_neuw
	br label %L1
L1:
	; If (cur < neuw)
	%r1 = load i64, i64* %P_cur
	%r2 = load i64, i64* %P_neuw
	%r3 = icmp slt i64 %r1, %r2
	br i1 %r3, label %L2, label %L3
L2:
	; True (cur < neuw)
	; return 1
	store i64 1, i64* %compare_retval
	%r4 = load i64, i64* %compare_retval
	ret i64 %r4
L3:
	; Else (cur < neuw)
	br label %L4
L4:
	; If (cur > neuw)
	%r5 = load i64, i64* %P_cur
	%r6 = load i64, i64* %P_neuw
	%r7 = icmp sgt i64 %r5, %r6
	br i1 %r7, label %L5, label %L6
L5:
	; True (cur > neuw)
	; return -1
	%r8 = sub i64 0, 1
	store i64 %r8, i64* %compare_retval
	%r9 = load i64, i64* %compare_retval
	ret i64 %r9
L6:
	; Else (cur > neuw)
	; return 0
	store i64 0, i64* %compare_retval
	%r10 = load i64, i64* %compare_retval
	ret i64 %r10
}

define void @addNode(i64 %numAdd, %struct.Node* %curr)
{
L7:
	%P_numAdd = alloca i64
	; parameters
	%P_curr = alloca %struct.Node*
	; local variables
	%compVal = alloca i64
	%newNode = alloca %struct.Node*
	store i64 %numAdd, i64* %P_numAdd
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L8
L8:
	; If (curr == nil)
	%r11 = load %struct.Node*, %struct.Node** %P_curr
	%r12 = load %struct.Node*, %struct.Node** @.nilNode
	%r13 = icmp eq %struct.Node* %r11, %r12
	br i1 %r13, label %L9, label %L10
L9:
	; True (curr == nil)
	; newNode = new Node;
	%r14 = call i8* @malloc(i32 24)
	%r15 = bitcast i8* %r14 to %struct.Node*
	store %struct.Node* %r15, %struct.Node** %newNode
	; newNode.value = numAdd;
	%r16 = load i64, i64* %P_numAdd
	%r17 = load %struct.Node*, %struct.Node** %newNode
	%r18 = getelementptr %struct.Node, %struct.Node* %r17, i32 0, i32 0
	store i64 %r16, i64* %r18
	; root = newNode;
	%r19 = load %struct.Node*, %struct.Node** %newNode
	store %struct.Node* %r19, %struct.Node** @root
	br label %L27
L10:
	; Else (curr == nil)
	; compVal = compare(curr.value, numAdd);
	%r20 = load %struct.Node*, %struct.Node** %P_curr
	%r21 = getelementptr %struct.Node, %struct.Node* %r20, i32 0, i32 0
	%r22 = load i64, i64* %r21
	%r23 = load i64, i64* %P_numAdd
	%r24 = call i64 @compare(i64 %r22, i64 %r23)
	store i64 %r24, i64* %compVal
	br label %L11
L11:
	; If (compVal == -1)
	%r25 = load i64, i64* %compVal
	%r26 = sub i64 0, 1
	%r27 = icmp eq i64 %r25, %r26
	br i1 %r27, label %L12, label %L13
L12:
	; True (compVal == -1)
	br label %L14
L14:
	; If (curr.left == nil)
	%r28 = load %struct.Node*, %struct.Node** %P_curr
	%r29 = getelementptr %struct.Node, %struct.Node* %r28, i32 0, i32 1
	%r30 = load %struct.Node*, %struct.Node** %r29
	%r31 = load %struct.Node*, %struct.Node** @.nilNode
	%r32 = icmp eq %struct.Node* %r30, %r31
	br i1 %r32, label %L15, label %L16
L15:
	; True (curr.left == nil)
	; newNode = new Node;
	%r33 = call i8* @malloc(i32 24)
	%r34 = bitcast i8* %r33 to %struct.Node*
	store %struct.Node* %r34, %struct.Node** %newNode
	; newNode.value = numAdd;
	%r35 = load i64, i64* %P_numAdd
	%r36 = load %struct.Node*, %struct.Node** %newNode
	%r37 = getelementptr %struct.Node, %struct.Node* %r36, i32 0, i32 0
	store i64 %r35, i64* %r37
	; curr.left = newNode;
	%r38 = load %struct.Node*, %struct.Node** %newNode
	%r39 = load %struct.Node*, %struct.Node** %P_curr
	%r40 = getelementptr %struct.Node, %struct.Node* %r39, i32 0, i32 1
	store %struct.Node* %r38, %struct.Node** %r40
	br label %L17
L16:
	; Else (curr.left == nil)
	%r41 = load i64, i64* %P_numAdd
	%r42 = load %struct.Node*, %struct.Node** %P_curr
	%r43 = getelementptr %struct.Node, %struct.Node* %r42, i32 0, i32 1
	%r44 = load %struct.Node*, %struct.Node** %r43
	call void @addNode(i64 %r41, %struct.Node* %r44)
	br label %L17
L17:
	; Exit (curr.left == nil)
	br label %L26
L13:
	; Else (compVal == -1)
	br label %L18
L18:
	; If (compVal == 1)
	%r45 = load i64, i64* %compVal
	%r46 = icmp eq i64 %r45, 1
	br i1 %r46, label %L19, label %L20
L19:
	; True (compVal == 1)
	br label %L21
L21:
	; If (curr.right == nil)
	%r47 = load %struct.Node*, %struct.Node** %P_curr
	%r48 = getelementptr %struct.Node, %struct.Node* %r47, i32 0, i32 2
	%r49 = load %struct.Node*, %struct.Node** %r48
	%r50 = load %struct.Node*, %struct.Node** @.nilNode
	%r51 = icmp eq %struct.Node* %r49, %r50
	br i1 %r51, label %L22, label %L23
L22:
	; True (curr.right == nil)
	; newNode = new Node;
	%r52 = call i8* @malloc(i32 24)
	%r53 = bitcast i8* %r52 to %struct.Node*
	store %struct.Node* %r53, %struct.Node** %newNode
	; newNode.value = numAdd;
	%r54 = load i64, i64* %P_numAdd
	%r55 = load %struct.Node*, %struct.Node** %newNode
	%r56 = getelementptr %struct.Node, %struct.Node* %r55, i32 0, i32 0
	store i64 %r54, i64* %r56
	; curr.right = newNode;
	%r57 = load %struct.Node*, %struct.Node** %newNode
	%r58 = load %struct.Node*, %struct.Node** %P_curr
	%r59 = getelementptr %struct.Node, %struct.Node* %r58, i32 0, i32 2
	store %struct.Node* %r57, %struct.Node** %r59
	br label %L24
L23:
	; Else (curr.right == nil)
	%r60 = load i64, i64* %P_numAdd
	%r61 = load %struct.Node*, %struct.Node** %P_curr
	%r62 = getelementptr %struct.Node, %struct.Node* %r61, i32 0, i32 2
	%r63 = load %struct.Node*, %struct.Node** %r62
	call void @addNode(i64 %r60, %struct.Node* %r63)
	br label %L24
L24:
	; Exit (curr.right == nil)
	br label %L25
L20:
	br label %L25
L25:
	; Exit (compVal == 1)
	br label %L26
L26:
	; Exit (compVal == -1)
	br label %L27
L27:
	; Exit (curr == nil)
	ret void
}

define void @printDepthTree(%struct.Node* %curr)
{
L28:
	%P_curr = alloca %struct.Node*
	; local variables
	%temp = alloca i64
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L29
L29:
	; If (curr != nil)
	%r64 = load %struct.Node*, %struct.Node** %P_curr
	%r65 = load %struct.Node*, %struct.Node** @.nilNode
	%r66 = icmp ne %struct.Node* %r64, %r65
	br i1 %r66, label %L30, label %L31
L30:
	; True (curr != nil)
	br label %L32
L32:
	; If (curr.left != nil)
	%r67 = load %struct.Node*, %struct.Node** %P_curr
	%r68 = getelementptr %struct.Node, %struct.Node* %r67, i32 0, i32 1
	%r69 = load %struct.Node*, %struct.Node** %r68
	%r70 = load %struct.Node*, %struct.Node** @.nilNode
	%r71 = icmp ne %struct.Node* %r69, %r70
	br i1 %r71, label %L33, label %L34
L33:
	; True (curr.left != nil)
	%r72 = load %struct.Node*, %struct.Node** %P_curr
	%r73 = getelementptr %struct.Node, %struct.Node* %r72, i32 0, i32 1
	%r74 = load %struct.Node*, %struct.Node** %r73
	call void @printDepthTree(%struct.Node* %r74)
	br label %L35
L34:
	br label %L35
L35:
	; Exit (curr.left != nil)
	; temp = curr.value;
	%r75 = load %struct.Node*, %struct.Node** %P_curr
	%r76 = getelementptr %struct.Node, %struct.Node* %r75, i32 0, i32 0
	%r77 = load i64, i64* %r76
	store i64 %r77, i64* %temp
	%r78 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r78)
	br label %L36
L36:
	; If (curr.right != nil)
	%r79 = load %struct.Node*, %struct.Node** %P_curr
	%r80 = getelementptr %struct.Node, %struct.Node* %r79, i32 0, i32 2
	%r81 = load %struct.Node*, %struct.Node** %r80
	%r82 = load %struct.Node*, %struct.Node** @.nilNode
	%r83 = icmp ne %struct.Node* %r81, %r82
	br i1 %r83, label %L37, label %L38
L37:
	; True (curr.right != nil)
	%r84 = load %struct.Node*, %struct.Node** %P_curr
	%r85 = getelementptr %struct.Node, %struct.Node* %r84, i32 0, i32 2
	%r86 = load %struct.Node*, %struct.Node** %r85
	call void @printDepthTree(%struct.Node* %r86)
	br label %L39
L38:
	br label %L39
L39:
	; Exit (curr.right != nil)
	br label %L40
L31:
	br label %L40
L40:
	; Exit (curr != nil)
	ret void
}

define void @deleteLeavesTree(%struct.Node* %curr)
{
L41:
	%P_curr = alloca %struct.Node*
	; local variables
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L42
L42:
	; If (curr != nil)
	%r87 = load %struct.Node*, %struct.Node** %P_curr
	%r88 = load %struct.Node*, %struct.Node** @.nilNode
	%r89 = icmp ne %struct.Node* %r87, %r88
	br i1 %r89, label %L43, label %L44
L43:
	; True (curr != nil)
	br label %L45
L45:
	; If (curr.left != nil)
	%r90 = load %struct.Node*, %struct.Node** %P_curr
	%r91 = getelementptr %struct.Node, %struct.Node* %r90, i32 0, i32 1
	%r92 = load %struct.Node*, %struct.Node** %r91
	%r93 = load %struct.Node*, %struct.Node** @.nilNode
	%r94 = icmp ne %struct.Node* %r92, %r93
	br i1 %r94, label %L46, label %L47
L46:
	; True (curr.left != nil)
	%r95 = load %struct.Node*, %struct.Node** %P_curr
	%r96 = getelementptr %struct.Node, %struct.Node* %r95, i32 0, i32 1
	%r97 = load %struct.Node*, %struct.Node** %r96
	call void @deleteLeavesTree(%struct.Node* %r97)
	br label %L48
L47:
	br label %L48
L48:
	; Exit (curr.left != nil)
	br label %L49
L49:
	; If (curr.right != nil)
	%r98 = load %struct.Node*, %struct.Node** %P_curr
	%r99 = getelementptr %struct.Node, %struct.Node* %r98, i32 0, i32 2
	%r100 = load %struct.Node*, %struct.Node** %r99
	%r101 = load %struct.Node*, %struct.Node** @.nilNode
	%r102 = icmp ne %struct.Node* %r100, %r101
	br i1 %r102, label %L50, label %L51
L50:
	; True (curr.right != nil)
	%r103 = load %struct.Node*, %struct.Node** %P_curr
	%r104 = getelementptr %struct.Node, %struct.Node* %r103, i32 0, i32 2
	%r105 = load %struct.Node*, %struct.Node** %r104
	call void @deleteLeavesTree(%struct.Node* %r105)
	br label %L52
L51:
	br label %L52
L52:
	; Exit (curr.right != nil)
	%r106 = load %struct.Node*, %struct.Node** %P_curr
	%r107 = bitcast %struct.Node* %r106 to i8*
	call void @free(i8* %r107)
	br label %L53
L44:
	br label %L53
L53:
	; Exit (curr != nil)
	ret void
}

define i64 @main()
{
L54:
	%_mainretval = alloca i64
	; local variables
	%input = alloca i64
	%temp = alloca i64
	; root = nil;
	%r108 = load %struct.Node*, %struct.Node** @.nilNode
	store %struct.Node* %r108, %struct.Node** @root
	; input = 0;
	store i64 0, i64* %input
	; input = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L55
L55:
	; for (input != 0)
	%r109 = load i64, i64* %input
	%r110 = icmp ne i64 %r109, 0
	br i1 %r110, label %L56, label %L57
L56:
	%r111 = load i64, i64* %input
	%r112 = load %struct.Node*, %struct.Node** @root
	call void @addNode(i64 %r111, %struct.Node* %r112)
	; input = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %input)
	br label %L55
L57:
	%r113 = load %struct.Node*, %struct.Node** @root
	call void @printDepthTree(%struct.Node* %r113)
	%r114 = load %struct.Node*, %struct.Node** @root
	call void @deleteLeavesTree(%struct.Node* %r114)
	store i64 0, i64* %_mainretval
	%r115 = load i64, i64* %_mainretval
	ret i64 %r115
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1