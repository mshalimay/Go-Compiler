source_filename = "Twiddleedee"
target triple = "x86_64-linux-gnu"

%struct.nums = type {i64, i64}
@.nilnums = common global %struct.nums* null


define i64 @fib1(i64 %f1)
{
L0:
	%fib1_retval = alloca i64
	%P_f1 = alloca i64
	; local variables
	store i64 %f1, i64* %P_f1
	br label %L1
L1:
	; If (f1 < 2)
	%r1 = load i64, i64* %P_f1
	%r2 = icmp slt i64 %r1, 2
	br i1 %r2, label %L2, label %L3
L2:
	; True (f1 < 2)
	; return f1
	%r3 = load i64, i64* %P_f1
	store i64 %r3, i64* %fib1_retval
	%r4 = load i64, i64* %fib1_retval
	ret i64 %r4
L3:
	; Else (f1 < 2)
	; return (fib1((f1 - 1)) + fib1((f1 - 2)))
	%r5 = load i64, i64* %P_f1
	%r6 = sub i64 %r5, 1
	%r7 = call i64 @fib1(i64 %r6)
	%r8 = load i64, i64* %P_f1
	%r9 = sub i64 %r8, 2
	%r10 = call i64 @fib1(i64 %r9)
	%r11 = add i64 %r7, %r10
	store i64 %r11, i64* %fib1_retval
	%r12 = load i64, i64* %fib1_retval
	ret i64 %r12
}

define i64 @fib2(i64 %f2)
{
L4:
	%fib2_retval = alloca i64
	%P_f2 = alloca i64
	; local variables
	%fir = alloca i64
	%sec = alloca i64
	%temp = alloca i64
	store i64 %f2, i64* %P_f2
	; fir = 0;
	store i64 0, i64* %fir
	; sec = 1;
	store i64 1, i64* %sec
	br label %L5
L5:
	; for (f2 != 0)
	%r13 = load i64, i64* %P_f2
	%r14 = icmp ne i64 %r13, 0
	br i1 %r14, label %L6, label %L7
L6:
	; f2 = (f2 - 1);
	%r15 = load i64, i64* %P_f2
	%r16 = sub i64 %r15, 1
	store i64 %r16, i64* %P_f2
	; temp = (fir + sec);
	%r17 = load i64, i64* %fir
	%r18 = load i64, i64* %sec
	%r19 = add i64 %r17, %r18
	store i64 %r19, i64* %temp
	; fir = sec;
	%r20 = load i64, i64* %sec
	store i64 %r20, i64* %fir
	; sec = temp;
	%r21 = load i64, i64* %temp
	store i64 %r21, i64* %sec
	br label %L5
L7:
	; return fir
	%r22 = load i64, i64* %fir
	store i64 %r22, i64* %fib2_retval
	%r23 = load i64, i64* %fib2_retval
	ret i64 %r23
}

define i64 @main()
{
L8:
	%_mainretval = alloca i64
	; local variables
	%x = alloca %struct.nums*
	%c = alloca i64
	%d = alloca i64
	%temp = alloca i64
	; x = new nums;
	%r24 = call i8* @malloc(i32 16)
	%r25 = bitcast i8* %r24 to %struct.nums*
	store %struct.nums* %r25, %struct.nums** %x
	; temp = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %temp)
	; x.a = temp;
	%r26 = load i64, i64* %temp
	%r27 = load %struct.nums*, %struct.nums** %x
	%r28 = getelementptr %struct.nums, %struct.nums* %r27, i32 0, i32 0
	store i64 %r26, i64* %r28
	; temp = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %temp)
	; x.b = temp;
	%r29 = load i64, i64* %temp
	%r30 = load %struct.nums*, %struct.nums** %x
	%r31 = getelementptr %struct.nums, %struct.nums* %r30, i32 0, i32 1
	store i64 %r29, i64* %r31
	; c = fib1(x.a);
	%r32 = load %struct.nums*, %struct.nums** %x
	%r33 = getelementptr %struct.nums, %struct.nums* %r32, i32 0, i32 0
	%r34 = load i64, i64* %r33
	%r35 = call i64 @fib1(i64 %r34)
	store i64 %r35, i64* %c
	; d = fib2(x.b);
	%r36 = load %struct.nums*, %struct.nums** %x
	%r37 = getelementptr %struct.nums, %struct.nums* %r36, i32 0, i32 1
	%r38 = load i64, i64* %r37
	%r39 = call i64 @fib2(i64 %r38)
	store i64 %r39, i64* %d
	%r40 = load %struct.nums*, %struct.nums** %x
	%r41 = bitcast %struct.nums* %r40 to i8*
	call void @free(i8* %r41)
	%r42 = load i64, i64* %c
	%r43 = load i64, i64* %d
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([12 x i8], [12 x i8]* @.fstr1, i32 0, i32 0), i64 %r42, i64 %r43)
	store i64 0, i64* %_mainretval
	%r44 = load i64, i64* %_mainretval
	ret i64 %r44
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [12 x i8] c"c=%ld\0Ad=%ld\00", align 1