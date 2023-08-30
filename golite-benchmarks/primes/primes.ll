source_filename = "primes"
target triple = "x86_64-linux-gnu"



define i64 @isqrt(i64 %a)
{
L0:
	%isqrt_retval = alloca i64
	%P_a = alloca i64
	; local variables
	%square = alloca i64
	%delta = alloca i64
	store i64 %a, i64* %P_a
	; square = 1;
	store i64 1, i64* %square
	; delta = 3;
	store i64 3, i64* %delta
	br label %L1
L1:
	; for (square <= a)
	%r1 = load i64, i64* %square
	%r2 = load i64, i64* %P_a
	%r3 = icmp sle i64 %r1, %r2
	br i1 %r3, label %L2, label %L3
L2:
	; square = (square + delta);
	%r4 = load i64, i64* %square
	%r5 = load i64, i64* %delta
	%r6 = add i64 %r4, %r5
	store i64 %r6, i64* %square
	; delta = (delta + 2);
	%r7 = load i64, i64* %delta
	%r8 = add i64 %r7, 2
	store i64 %r8, i64* %delta
	br label %L1
L3:
	; return (((delta / 2) - 1))
	%r9 = load i64, i64* %delta
	%r10 = sdiv i64 %r9, 2
	%r11 = sub i64 %r10, 1
	store i64 %r11, i64* %isqrt_retval
	%r12 = load i64, i64* %isqrt_retval
	ret i64 %r12
}

define i64 @prime(i64 %a)
{
L4:
	%prime_retval = alloca i64
	%P_a = alloca i64
	; local variables
	%max = alloca i64
	%divisor = alloca i64
	%remainder = alloca i64
	store i64 %a, i64* %P_a
	br label %L5
L5:
	; If (a < 2)
	%r13 = load i64, i64* %P_a
	%r14 = icmp slt i64 %r13, 2
	br i1 %r14, label %L6, label %L7
L6:
	; True (a < 2)
	; return false
	store i64 0, i64* %prime_retval
	%r15 = load i64, i64* %prime_retval
	ret i64 %r15
L7:
	; Else (a < 2)
	; max = isqrt(a);
	%r16 = load i64, i64* %P_a
	%r17 = call i64 @isqrt(i64 %r16)
	store i64 %r17, i64* %max
	; divisor = 2;
	store i64 2, i64* %divisor
	br label %L8
L8:
	; for (divisor <= max)
	%r18 = load i64, i64* %divisor
	%r19 = load i64, i64* %max
	%r20 = icmp sle i64 %r18, %r19
	br i1 %r20, label %L9, label %L14
L9:
	; remainder = (a - ((((a / divisor)) * divisor)));
	%r21 = load i64, i64* %P_a
	%r22 = load i64, i64* %P_a
	%r23 = load i64, i64* %divisor
	%r24 = sdiv i64 %r22, %r23
	%r25 = load i64, i64* %divisor
	%r26 = mul i64 %r24, %r25
	%r27 = sub i64 %r21, %r26
	store i64 %r27, i64* %remainder
	br label %L10
L10:
	; If (remainder == 0)
	%r28 = load i64, i64* %remainder
	%r29 = icmp eq i64 %r28, 0
	br i1 %r29, label %L11, label %L12
L11:
	; True (remainder == 0)
	; return false
	store i64 0, i64* %prime_retval
	%r30 = load i64, i64* %prime_retval
	ret i64 %r30
L12:
	br label %L13
L13:
	; Exit (remainder == 0)
	; divisor = (divisor + 1);
	%r31 = load i64, i64* %divisor
	%r32 = add i64 %r31, 1
	store i64 %r32, i64* %divisor
	br label %L8
L14:
	; return true
	store i64 1, i64* %prime_retval
	%r33 = load i64, i64* %prime_retval
	ret i64 %r33
}

define i64 @main()
{
L15:
	%_mainretval = alloca i64
	; local variables
	%limit = alloca i64
	%a = alloca i64
	; limit = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %limit)
	; a = 0;
	store i64 0, i64* %a
	br label %L16
L16:
	; for (a <= limit)
	%r34 = load i64, i64* %a
	%r35 = load i64, i64* %limit
	%r36 = icmp sle i64 %r34, %r35
	br i1 %r36, label %L17, label %L22
L17:
	br label %L18
L18:
	; If prime(a)
	%r37 = load i64, i64* %a
	%r38 = call i64 @prime(i64 %r37)
	%r39 = trunc i64 %r38 to i1
	br i1 %r39, label %L19, label %L20
L19:
	; True prime(a)
	%r40 = load i64, i64* %a
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.fstr1, i32 0, i32 0), i64 %r40)
	br label %L21
L20:
	br label %L21
L21:
	; Exit prime(a)
	; a = (a + 1);
	%r41 = load i64, i64* %a
	%r42 = add i64 %r41, 1
	store i64 %r42, i64* %a
	br label %L16
L22:
	store i64 0, i64* %_mainretval
	%r43 = load i64, i64* %_mainretval
	ret i64 %r43
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1