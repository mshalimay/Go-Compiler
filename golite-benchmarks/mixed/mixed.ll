source_filename = "mixed"
target triple = "x86_64-linux-gnu"

%struct.simple = type {i64}
@.nilsimple = common global %struct.simple* null
%struct.foo = type {i64, i64, %struct.simple*}
@.nilfoo = common global %struct.foo* null

@globalfoo = common global %struct.foo* null
@unusedGlobal = common global %struct.foo* null

define void @tailrecursive(i64 %num)
{
L0:
	%P_num = alloca i64
	; local variables
	%unused = alloca %struct.foo*
	store i64 %num, i64* %P_num
	br label %L1
L1:
	; If (num <= 0)
	%r1 = load i64, i64* %P_num
	%r2 = icmp sle i64 %r1, 0
	br i1 %r2, label %L2, label %L3
L2:
	; True (num <= 0)
	; return
	ret void
L3:
	br label %L4
L4:
	; Exit (num <= 0)
	; unused = new foo;
	%r3 = call i8* @malloc(i32 24)
	%r4 = bitcast i8* %r3 to %struct.foo*
	store %struct.foo* %r4, %struct.foo** %unused
	; unusedGlobal = unused;
	%r5 = load %struct.foo*, %struct.foo** %unused
	store %struct.foo* %r5, %struct.foo** @unusedGlobal
	%r6 = load i64, i64* %P_num
	%r7 = sub i64 %r6, 1
	call void @tailrecursive(i64 %r7)
	ret void
}

define i64 @add(i64 %x, i64 %y)
{
L5:
	%add_retval = alloca i64
	%P_x = alloca i64
	; parameters
	%P_y = alloca i64
	; local variables
	store i64 %x, i64* %P_x
	store i64 %y, i64* %P_y
	; return (x + y)
	%r8 = load i64, i64* %P_x
	%r9 = load i64, i64* %P_y
	%r10 = add i64 %r8, %r9
	store i64 %r10, i64* %add_retval
	%r11 = load i64, i64* %add_retval
	ret i64 %r11
}

define void @domath(i64 %num)
{
L6:
	%P_num = alloca i64
	; local variables
	%math1 = alloca %struct.foo*
	%math2 = alloca %struct.foo*
	%tmp = alloca i64
	store i64 %num, i64* %P_num
	; math1 = new foo;
	%r12 = call i8* @malloc(i32 24)
	%r13 = bitcast i8* %r12 to %struct.foo*
	store %struct.foo* %r13, %struct.foo** %math1
	; math1.simp = new simple;
	%r14 = call i8* @malloc(i32 8)
	%r15 = bitcast i8* %r14 to %struct.simple*
	%r16 = load %struct.foo*, %struct.foo** %math1
	%r17 = getelementptr %struct.foo, %struct.foo* %r16, i32 0, i32 2
	store %struct.simple* %r15, %struct.simple** %r17
	; math2 = new foo;
	%r18 = call i8* @malloc(i32 24)
	%r19 = bitcast i8* %r18 to %struct.foo*
	store %struct.foo* %r19, %struct.foo** %math2
	; math2.simp = new simple;
	%r20 = call i8* @malloc(i32 8)
	%r21 = bitcast i8* %r20 to %struct.simple*
	%r22 = load %struct.foo*, %struct.foo** %math2
	%r23 = getelementptr %struct.foo, %struct.foo* %r22, i32 0, i32 2
	store %struct.simple* %r21, %struct.simple** %r23
	; math1.bar = num;
	%r24 = load i64, i64* %P_num
	%r25 = load %struct.foo*, %struct.foo** %math1
	%r26 = getelementptr %struct.foo, %struct.foo* %r25, i32 0, i32 0
	store i64 %r24, i64* %r26
	; math2.bar = 3;
	%r27 = load %struct.foo*, %struct.foo** %math2
	%r28 = getelementptr %struct.foo, %struct.foo* %r27, i32 0, i32 0
	store i64 3, i64* %r28
	; math1.simp.one = math1.bar;
	%r29 = load %struct.foo*, %struct.foo** %math1
	%r30 = getelementptr %struct.foo, %struct.foo* %r29, i32 0, i32 0
	%r31 = load i64, i64* %r30
	%r32 = load %struct.foo*, %struct.foo** %math1
	%r33 = getelementptr %struct.foo, %struct.foo* %r32, i32 0, i32 2
	%r34 = load %struct.simple*, %struct.simple** %r33
	%r35 = getelementptr %struct.simple, %struct.simple* %r34, i32 0, i32 0
	store i64 %r31, i64* %r35
	; math2.simp.one = math2.bar;
	%r36 = load %struct.foo*, %struct.foo** %math2
	%r37 = getelementptr %struct.foo, %struct.foo* %r36, i32 0, i32 0
	%r38 = load i64, i64* %r37
	%r39 = load %struct.foo*, %struct.foo** %math2
	%r40 = getelementptr %struct.foo, %struct.foo* %r39, i32 0, i32 2
	%r41 = load %struct.simple*, %struct.simple** %r40
	%r42 = getelementptr %struct.simple, %struct.simple* %r41, i32 0, i32 0
	store i64 %r38, i64* %r42
	br label %L7
L7:
	; for (num > 0)
	%r43 = load i64, i64* %P_num
	%r44 = icmp sgt i64 %r43, 0
	br i1 %r44, label %L8, label %L9
L8:
	; tmp = (math1.bar * math2.bar);
	%r45 = load %struct.foo*, %struct.foo** %math1
	%r46 = getelementptr %struct.foo, %struct.foo* %r45, i32 0, i32 0
	%r47 = load i64, i64* %r46
	%r48 = load %struct.foo*, %struct.foo** %math2
	%r49 = getelementptr %struct.foo, %struct.foo* %r48, i32 0, i32 0
	%r50 = load i64, i64* %r49
	%r51 = mul i64 %r47, %r50
	store i64 %r51, i64* %tmp
	; tmp = (((tmp * math1.simp.one)) / math2.bar);
	%r52 = load i64, i64* %tmp
	%r53 = load %struct.foo*, %struct.foo** %math1
	%r54 = getelementptr %struct.foo, %struct.foo* %r53, i32 0, i32 2
	%r55 = load %struct.simple*, %struct.simple** %r54
	%r56 = getelementptr %struct.simple, %struct.simple* %r55, i32 0, i32 0
	%r57 = load i64, i64* %r56
	%r58 = mul i64 %r52, %r57
	%r59 = load %struct.foo*, %struct.foo** %math2
	%r60 = getelementptr %struct.foo, %struct.foo* %r59, i32 0, i32 0
	%r61 = load i64, i64* %r60
	%r62 = sdiv i64 %r58, %r61
	store i64 %r62, i64* %tmp
	; tmp = add(math2.simp.one, math1.bar);
	%r63 = load %struct.foo*, %struct.foo** %math2
	%r64 = getelementptr %struct.foo, %struct.foo* %r63, i32 0, i32 2
	%r65 = load %struct.simple*, %struct.simple** %r64
	%r66 = getelementptr %struct.simple, %struct.simple* %r65, i32 0, i32 0
	%r67 = load i64, i64* %r66
	%r68 = load %struct.foo*, %struct.foo** %math1
	%r69 = getelementptr %struct.foo, %struct.foo* %r68, i32 0, i32 0
	%r70 = load i64, i64* %r69
	%r71 = call i64 @add(i64 %r67, i64 %r70)
	store i64 %r71, i64* %tmp
	; tmp = (math2.bar - math1.bar);
	%r72 = load %struct.foo*, %struct.foo** %math2
	%r73 = getelementptr %struct.foo, %struct.foo* %r72, i32 0, i32 0
	%r74 = load i64, i64* %r73
	%r75 = load %struct.foo*, %struct.foo** %math1
	%r76 = getelementptr %struct.foo, %struct.foo* %r75, i32 0, i32 0
	%r77 = load i64, i64* %r76
	%r78 = sub i64 %r74, %r77
	store i64 %r78, i64* %tmp
	; num = (num - 1);
	%r79 = load i64, i64* %P_num
	%r80 = sub i64 %r79, 1
	store i64 %r80, i64* %P_num
	br label %L7
L9:
	%r81 = load %struct.foo*, %struct.foo** %math1
	%r82 = bitcast %struct.foo* %r81 to i8*
	call void @free(i8* %r82)
	%r83 = load %struct.foo*, %struct.foo** %math2
	%r84 = bitcast %struct.foo* %r83 to i8*
	call void @free(i8* %r84)
	ret void
}

define void @objinstantiation(i64 %num)
{
L10:
	%P_num = alloca i64
	; local variables
	%tmp = alloca %struct.foo*
	store i64 %num, i64* %P_num
	br label %L11
L11:
	; for (num > 0)
	%r85 = load i64, i64* %P_num
	%r86 = icmp sgt i64 %r85, 0
	br i1 %r86, label %L12, label %L13
L12:
	; tmp = new foo;
	%r87 = call i8* @malloc(i32 24)
	%r88 = bitcast i8* %r87 to %struct.foo*
	store %struct.foo* %r88, %struct.foo** %tmp
	%r89 = load %struct.foo*, %struct.foo** %tmp
	%r90 = bitcast %struct.foo* %r89 to i8*
	call void @free(i8* %r90)
	; num = (num - 1);
	%r91 = load i64, i64* %P_num
	%r92 = sub i64 %r91, 1
	store i64 %r92, i64* %P_num
	br label %L11
L13:
	ret void
}

define i64 @ackermann(i64 %m, i64 %n)
{
L14:
	%ackermann_retval = alloca i64
	%P_m = alloca i64
	; parameters
	%P_n = alloca i64
	; local variables
	store i64 %m, i64* %P_m
	store i64 %n, i64* %P_n
	br label %L15
L15:
	; If (m == 0)
	%r93 = load i64, i64* %P_m
	%r94 = icmp eq i64 %r93, 0
	br i1 %r94, label %L16, label %L17
L16:
	; True (m == 0)
	; return (n + 1)
	%r95 = load i64, i64* %P_n
	%r96 = add i64 %r95, 1
	store i64 %r96, i64* %ackermann_retval
	%r97 = load i64, i64* %ackermann_retval
	ret i64 %r97
L17:
	br label %L18
L18:
	; Exit (m == 0)
	br label %L19
L19:
	; If (n == 0)
	%r98 = load i64, i64* %P_n
	%r99 = icmp eq i64 %r98, 0
	br i1 %r99, label %L20, label %L21
L20:
	; True (n == 0)
	; return ackermann((m - 1), 1)
	%r100 = load i64, i64* %P_m
	%r101 = sub i64 %r100, 1
	%r102 = call i64 @ackermann(i64 %r101, i64 1)
	store i64 %r102, i64* %ackermann_retval
	%r103 = load i64, i64* %ackermann_retval
	ret i64 %r103
L21:
	; Else (n == 0)
	; return ackermann((m - 1), ackermann(m, (n - 1)))
	%r104 = load i64, i64* %P_m
	%r105 = sub i64 %r104, 1
	%r106 = load i64, i64* %P_m
	%r107 = load i64, i64* %P_n
	%r108 = sub i64 %r107, 1
	%r109 = call i64 @ackermann(i64 %r106, i64 %r108)
	%r110 = call i64 @ackermann(i64 %r105, i64 %r109)
	store i64 %r110, i64* %ackermann_retval
	%r111 = load i64, i64* %ackermann_retval
	ret i64 %r111
}

define i64 @main()
{
L22:
	%_mainretval = alloca i64
	; local variables
	%a = alloca i64
	%b = alloca i64
	%c = alloca i64
	%d = alloca i64
	%e = alloca i64
	%temp = alloca i64
	; a = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %a)
	; b = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %b)
	; c = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %c)
	; d = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %d)
	; e = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %e)
	%r112 = load i64, i64* %a
	call void @tailrecursive(i64 %r112)
	%r113 = load i64, i64* %b
	call void @domath(i64 %r113)
	%r114 = load i64, i64* %c
	call void @objinstantiation(i64 %r114)
	; temp = ackermann(d, e);
	%r115 = load i64, i64* %d
	%r116 = load i64, i64* %e
	%r117 = call i64 @ackermann(i64 %r115, i64 %r116)
	store i64 %r117, i64* %temp
	%r118 = load i64, i64* %a
	%r119 = load i64, i64* %b
	%r120 = load i64, i64* %c
	%r121 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([27 x i8], [27 x i8]* @.fstr1, i32 0, i32 0), i64 %r118, i64 %r119, i64 %r120, i64 %r121)
	store i64 0, i64* %_mainretval
	%r122 = load i64, i64* %_mainretval
	ret i64 %r122
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [27 x i8] c"a=%ld\0Ab=%ld\0Ac=%ld,temp=%ld\00", align 1