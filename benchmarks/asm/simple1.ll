source_filename = "simple1"
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
	%r0 = load i64, i64* %P_x
	%r1 = icmp sle i64 %r0, 1
	br i1 %r1, label %L2, label %L3
L2:
	; True (x <= 1)
	; return 1
	store i64 1, i64* %fact_retval
	%r2 = load i64, i64* %fact_retval
	ret i64 %r2
L3:
	; Else (x <= 1)
	; return ((x * fact((x - 1))))
	%r3 = load i64, i64* %P_x
	%r4 = load i64, i64* %P_x
	%r5 = sub i64 %r4, 1
	%r6 = call i64 @fact(i64 %r5)
	%r7 = mul i64 %r3, %r6
	store i64 %r7, i64* %fact_retval
	%r8 = load i64, i64* %fact_retval
	ret i64 %r8
}

define i64 @main()
{
L4:
	%main_retval = alloca i64
	; local variables
	%factor = alloca i64
	%toStop = alloca i64
	%temp = alloca i64
	%stop = alloca i64
	; stop = false;
	store i64 0, i64* %stop
	; factor = 0;
	store i64 0, i64* %factor
	; factor = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %factor)
	; temp = fact(factor);
	%r9 = load i64, i64* %factor
	%r10 = call i64 @fact(i64 %r9)
	store i64 %r10, i64* %temp
	%r11 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r11)
	store i64 0, i64* %main_retval
	%r12 = load i64, i64* %main_retval
	ret i64 %r12
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1