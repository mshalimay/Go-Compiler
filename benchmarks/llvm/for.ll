source_filename = "for"
target triple = "x86_64-linux-gnu"



define i64 @main()
{
L0:
	%_mainretval = alloca i64
	%x = alloca i64
	%r1 = sub i64 0, 10
	store i64 %r1, i64* %x
	br label %L1
L1:
	%r2 = load i64, i64* %x
	%r3 = icmp sgt i64 %r2, 0
	br i1 %r3, label %L2, label %L3
L2:
	%r4 = load i64, i64* %x
	%r5 = add i64 %r4, 1
	store i64 %r5, i64* %x
	br label %L1
L3:
	store i64 0, i64* %_mainretval
	%r6 = load i64, i64* %_mainretval
	ret i64 %r6
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)
