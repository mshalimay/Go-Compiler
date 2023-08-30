source_filename = "nested"
target triple = "x86_64-linux-gnu"



define i64 @main()
{
L0:
	%_mainretval = alloca i64
	; memory alloc local variables
	%x = alloca i64
	; x = -10;
	%r1 = sub i64 0, 10
	store i64 %r1, i64* %x
	br label %L1
L1:
	; If (x < 0)
	%r2 = load i64, i64* %x
	%r3 = icmp slt i64 %r2, 0
	br i1 %r3, label %L2, label %L3
L2:
	; x = 0;
	store i64 0, i64* %x
	br label %L4
L4:
	; If (x == 0)
	%r4 = load i64, i64* %x
	%r5 = icmp eq i64 %r4, 0
	br i1 %r5, label %L5, label %L6
L5:
	; x = 1;
	store i64 1, i64* %x
	br label %L7
L6:
	br label %L7
L7:
	br label %L8
L3:
	; Else (x < 0)
	; x = 1;
	store i64 1, i64* %x
	store i64 0, i64* %_mainretval
	%r6 = load i64, i64* %_mainretval
	ret i64 %r6
L8:
	store i64 0, i64* %_mainretval
	%r7 = load i64, i64* %_mainretval
	ret i64 %r7
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
