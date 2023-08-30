source_filename = "thermopylae"
target triple = "x86_64-linux-gnu"



define i64 @fact(i64 %x)
{
L0:
	%fact_retval = alloca i64
	%P_x = alloca i64
	; local variables
	store i64 %x, i64* %P_x
	br label %L1
L1:
	; If (x <= 1)
	%r1 = load i64, i64* %P_x
	%r2 = icmp sle i64 %r1, 1
	br i1 %r2, label %L2, label %L3
L2:
	; True (x <= 1)
	; return 1
	store i64 1, i64* %fact_retval
	%r3 = load i64, i64* %fact_retval
	ret i64 %r3
L3:
	; Else (x <= 1)
	; return ((x * fact((x - 1))))
	%r4 = load i64, i64* %P_x
	%r5 = load i64, i64* %P_x
	%r6 = sub i64 %r5, 1
	%r7 = call i64 @fact(i64 %r6)
	%r8 = mul i64 %r4, %r7
	store i64 %r8, i64* %fact_retval
	%r9 = load i64, i64* %fact_retval
	ret i64 %r9
}

define i64 @main()
{
L4:
	%main_retval = alloca i64
	; local variables
	%stop = alloca i64
	%factor = alloca i64
	%toStop = alloca i64
	%temp = alloca i64
	; stop = false;
	store i64 0, i64* %stop
	; factor = 0;
	store i64 0, i64* %factor
	br label %L5
L5:
	; for !stop
	%r10 = load i64, i64* %stop
	%r11 = xor i64 %r10, 1
	%r12 = trunc i64 %r11 to i1
	br i1 %r12, label %L6, label %L11
L6:
	; factor = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %factor)
	; temp = fact(factor);
	%r13 = load i64, i64* %factor
	%r14 = call i64 @fact(i64 %r13)
	store i64 %r14, i64* %temp
	%r15 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r15)
	; toStop = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %toStop)
	br label %L7
L7:
	; If (toStop == 0)
	%r16 = load i64, i64* %toStop
	%r17 = icmp eq i64 %r16, 0
	br i1 %r17, label %L8, label %L9
L8:
	; True (toStop == 0)
	; stop = true;
	store i64 1, i64* %stop
	br label %L10
L9:
	br label %L10
L10:
	; Exit (toStop == 0)
	br label %L5
L11:
	store i64 0, i64* %main_retval
	%r18 = load i64, i64* %main_retval
	ret i64 %r18
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1