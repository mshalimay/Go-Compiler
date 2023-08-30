source_filename = "linkedlist"
target triple = "x86_64-linux-gnu"

%struct.Node = type {i64, %struct.Node*}
@.nilNode = common global %struct.Node* null
%struct.Node2 = type {i64, %struct.Node*}
@.nilNode2 = common global %struct.Node2* null

@head = common global %struct.Node* null
@tail = common global %struct.Node* null

define void @Add(i64 %num)
{
L0:
	%P_num = alloca i64
	; local variables
	%newList = alloca %struct.Node*
	store i64 %num, i64* %P_num
	; newList = new Node;
	%r0 = call i8* @malloc(i32 16)
	%r1 = bitcast i8* %r0 to %struct.Node*
	store %struct.Node* %r1, %struct.Node** %newList
	; newList.num = num;
	%r2 = load i64, i64* %P_num
	%r3 = load %struct.Node*, %struct.Node** %newList
	%r4 = getelementptr %struct.Node, %struct.Node* %r3, i32 0, i32 0
	store i64 %r2, i64* %r4
	; newList.next = nil;
	%r5 = load %struct.Node*, %struct.Node** @.nilNode
	%r6 = load %struct.Node*, %struct.Node** %newList
	%r7 = getelementptr %struct.Node, %struct.Node* %r6, i32 0, i32 1
	store %struct.Node* %r5, %struct.Node** %r7
	br label %L1
L1:
	; If (head == nil)
	%r8 = load %struct.Node*, %struct.Node** @head
	%r9 = load %struct.Node*, %struct.Node** @.nilNode
	%r10 = icmp eq %struct.Node* %r8, %r9
	br i1 %r10, label %L2, label %L3
L2:
	; True (head == nil)
	; head = newList;
	%r11 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r11, %struct.Node** @head
	; tail = newList;
	%r12 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r12, %struct.Node** @tail
	br label %L4
L3:
	; Else (head == nil)
	; tail.next = newList;
	%r13 = load %struct.Node*, %struct.Node** %newList
	%r14 = load %struct.Node*, %struct.Node** @tail
	%r15 = getelementptr %struct.Node, %struct.Node* %r14, i32 0, i32 1
	store %struct.Node* %r13, %struct.Node** %r15
	; tail = newList;
	%r16 = load %struct.Node*, %struct.Node** %newList
	store %struct.Node* %r16, %struct.Node** @tail
	br label %L4
L4:
	; Exit (head == nil)
	ret void
}

define void @PrintList(%struct.Node* %curr)
{
L5:
	%P_curr = alloca %struct.Node*
	; local variables
	%printValue = alloca i64
	store %struct.Node* %curr, %struct.Node** %P_curr
	br label %L6
L6:
	; If (curr == tail)
	%r17 = load %struct.Node*, %struct.Node** %P_curr
	%r18 = load %struct.Node*, %struct.Node** @tail
	%r19 = icmp eq %struct.Node* %r17, %r18
	br i1 %r19, label %L7, label %L8
L7:
	; True (curr == tail)
	; printValue = curr.num;
	%r20 = load %struct.Node*, %struct.Node** %P_curr
	%r21 = getelementptr %struct.Node, %struct.Node* %r20, i32 0, i32 0
	%r22 = load i64, i64* %r21
	store i64 %r22, i64* %printValue
	%r23 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r23)
	br label %L9
L8:
	; Else (curr == tail)
	; printValue = curr.num;
	%r24 = load %struct.Node*, %struct.Node** %P_curr
	%r25 = getelementptr %struct.Node, %struct.Node* %r24, i32 0, i32 0
	%r26 = load i64, i64* %r25
	store i64 %r26, i64* %printValue
	%r27 = load i64, i64* %printValue
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr2, i32 0, i32 0), i64 %r27)
	%r28 = load %struct.Node*, %struct.Node** %P_curr
	%r29 = getelementptr %struct.Node, %struct.Node* %r28, i32 0, i32 1
	%r30 = load %struct.Node*, %struct.Node** %r29
	call void @PrintList(%struct.Node* %r30)
	br label %L9
L9:
	; Exit (curr == tail)
	ret void
}

define void @Del(%struct.Node* %curr, i64 %num)
{
L10:
	%P_curr = alloca %struct.Node*
	; parameters
	%P_num = alloca i64
	; local variables
	%temp = alloca %struct.Node*
	store %struct.Node* %curr, %struct.Node** %P_curr
	store i64 %num, i64* %P_num
	br label %L11
L11:
	; If (curr == nil)
	%r31 = load %struct.Node*, %struct.Node** %P_curr
	%r32 = load %struct.Node*, %struct.Node** @.nilNode
	%r33 = icmp eq %struct.Node* %r31, %r32
	br i1 %r33, label %L12, label %L13
L12:
	; True (curr == nil)
	br label %L26
L13:
	; Else (curr == nil)
	br label %L14
L14:
	; If (head.num == num)
	%r34 = load %struct.Node*, %struct.Node** @head
	%r35 = getelementptr %struct.Node, %struct.Node* %r34, i32 0, i32 0
	%r36 = load i64, i64* %r35
	%r37 = load i64, i64* %P_num
	%r38 = icmp eq i64 %r36, %r37
	br i1 %r38, label %L15, label %L16
L15:
	; True (head.num == num)
	; temp = head;
	%r39 = load %struct.Node*, %struct.Node** @head
	store %struct.Node* %r39, %struct.Node** %temp
	; head = head.next;
	%r40 = load %struct.Node*, %struct.Node** @head
	%r41 = getelementptr %struct.Node, %struct.Node* %r40, i32 0, i32 1
	%r42 = load %struct.Node*, %struct.Node** %r41
	store %struct.Node* %r42, %struct.Node** @head
	%r43 = load %struct.Node*, %struct.Node** %temp
	%r44 = bitcast %struct.Node* %r43 to i8*
	call void @free(i8* %r44)
	br label %L25
L16:
	; Else (head.num == num)
	br label %L17
L17:
	; If (curr.next == tail)
	%r45 = load %struct.Node*, %struct.Node** %P_curr
	%r46 = getelementptr %struct.Node, %struct.Node* %r45, i32 0, i32 1
	%r47 = load %struct.Node*, %struct.Node** %r46
	%r48 = load %struct.Node*, %struct.Node** @tail
	%r49 = icmp eq %struct.Node* %r47, %r48
	br i1 %r49, label %L18, label %L19
L18:
	; True (curr.next == tail)
	; temp = tail;
	%r50 = load %struct.Node*, %struct.Node** @tail
	store %struct.Node* %r50, %struct.Node** %temp
	; tail = curr;
	%r51 = load %struct.Node*, %struct.Node** %P_curr
	store %struct.Node* %r51, %struct.Node** @tail
	; tail.next = nil;
	%r52 = load %struct.Node*, %struct.Node** @.nilNode
	%r53 = load %struct.Node*, %struct.Node** @tail
	%r54 = getelementptr %struct.Node, %struct.Node* %r53, i32 0, i32 1
	store %struct.Node* %r52, %struct.Node** %r54
	%r55 = load %struct.Node*, %struct.Node** %temp
	%r56 = bitcast %struct.Node* %r55 to i8*
	call void @free(i8* %r56)
	br label %L24
L19:
	; Else (curr.next == tail)
	br label %L20
L20:
	; If ((curr.next.num == num))
	%r57 = load %struct.Node*, %struct.Node** %P_curr
	%r58 = getelementptr %struct.Node, %struct.Node* %r57, i32 0, i32 1
	%r59 = load %struct.Node*, %struct.Node** %r58
	%r60 = getelementptr %struct.Node, %struct.Node* %r59, i32 0, i32 0
	%r61 = load i64, i64* %r60
	%r62 = load i64, i64* %P_num
	%r63 = icmp eq i64 %r61, %r62
	br i1 %r63, label %L21, label %L22
L21:
	; True ((curr.next.num == num))
	; temp = curr.next;
	%r64 = load %struct.Node*, %struct.Node** %P_curr
	%r65 = getelementptr %struct.Node, %struct.Node* %r64, i32 0, i32 1
	%r66 = load %struct.Node*, %struct.Node** %r65
	store %struct.Node* %r66, %struct.Node** %temp
	; curr.next = curr.next.next;
	%r67 = load %struct.Node*, %struct.Node** %P_curr
	%r68 = getelementptr %struct.Node, %struct.Node* %r67, i32 0, i32 1
	%r69 = load %struct.Node*, %struct.Node** %r68
	%r70 = getelementptr %struct.Node, %struct.Node* %r69, i32 0, i32 1
	%r71 = load %struct.Node*, %struct.Node** %r70
	%r72 = load %struct.Node*, %struct.Node** %P_curr
	%r73 = getelementptr %struct.Node, %struct.Node* %r72, i32 0, i32 1
	store %struct.Node* %r71, %struct.Node** %r73
	%r74 = load %struct.Node*, %struct.Node** %temp
	%r75 = bitcast %struct.Node* %r74 to i8*
	call void @free(i8* %r75)
	br label %L23
L22:
	; Else ((curr.next.num == num))
	%r76 = load %struct.Node*, %struct.Node** %P_curr
	%r77 = getelementptr %struct.Node, %struct.Node* %r76, i32 0, i32 1
	%r78 = load %struct.Node*, %struct.Node** %r77
	%r79 = load i64, i64* %P_num
	call void @Del(%struct.Node* %r78, i64 %r79)
	br label %L23
L23:
	; Exit ((curr.next.num == num))
	br label %L24
L24:
	; Exit (curr.next == tail)
	br label %L25
L25:
	; Exit (head.num == num)
	br label %L26
L26:
	; Exit (curr == nil)
	ret void
}

define i64 @main()
{
L27:
	%main_retval = alloca i64
	; local variables
	%x = alloca i64
	%y = alloca i64
	%i = alloca i64
	; x = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %x)
	; y = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %y)
	call void @Add(i64 1)
	call void @Add(i64 10)
	call void @Add(i64 3)
	call void @Add(i64 4)
	%r80 = load i64, i64* %x
	call void @Add(i64 %r80)
	%r81 = load %struct.Node*, %struct.Node** @head
	call void @PrintList(%struct.Node* %r81)
	; i = 0;
	store i64 0, i64* %i
	br label %L28
L28:
	; for ((i < 50000000))
	%r82 = load i64, i64* %i
	%r83 = icmp slt i64 %r82, 50000000
	br i1 %r83, label %L29, label %L30
L29:
	%r84 = load i64, i64* %i
	call void @Add(i64 %r84)
	; i = ((i + 1));
	%r85 = load i64, i64* %i
	%r86 = add i64 %r85, 1
	store i64 %r86, i64* %i
	br label %L28
L30:
	; i = 0;
	store i64 0, i64* %i
	br label %L31
L31:
	; for ((i < 50000000))
	%r87 = load i64, i64* %i
	%r88 = icmp slt i64 %r87, 50000000
	br i1 %r88, label %L32, label %L33
L32:
	%r89 = load %struct.Node*, %struct.Node** @head
	%r90 = load i64, i64* %i
	call void @Del(%struct.Node* %r89, i64 %r90)
	; i = ((i + 1));
	%r91 = load i64, i64* %i
	%r92 = add i64 %r91, 1
	store i64 %r92, i64* %i
	br label %L31
L33:
	%r93 = load %struct.Node*, %struct.Node** @head
	%r94 = load i64, i64* %y
	call void @Del(%struct.Node* %r93, i64 %r94)
	%r95 = load %struct.Node*, %struct.Node** @head
	call void @PrintList(%struct.Node* %r95)
	store i64 0, i64* %main_retval
	%r96 = load i64, i64* %main_retval
	ret i64 %r96
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr2 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1