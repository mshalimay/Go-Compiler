source_filename = "twoifs"
target triple = "/home/mshalimay/Compilers/proj-mashalimay/benchmarks/llvm/twoifs.golite"


@Node = common global i64 0

define void @printDepthTree(i64 %curr)
{
L0:
	%P_curr = alloca i64
	; local variables
	%temp = alloca i64
	store i64 %curr, i64* %P_curr
	br label %L1
L1:
	; If (curr != 0)
	%r1 = load i64, i64* %P_curr
	%r2 = icmp ne i64 %r1, 0
	br i1 %r2, label %L2, label %L3
L2:
	; True (curr != 0)
	; curr = 0;
	store i64 0, i64* %P_curr
	br label %L4
L4:
	; If (curr > 0)
	%r3 = load i64, i64* %P_curr
	%r4 = icmp sgt i64 %r3, 0
	br i1 %r4, label %L5, label %L6
L5:
	; True (curr > 0)
	%r5 = load i64, i64* %P_curr
	call void @printDepthTree(i64 %r5)
	br label %L7
L6:
	br label %L7
L7:
	; Exit (curr > 0)
	; temp = Node;
	%r6 = load i64, i64* @Node
	store i64 %r6, i64* %temp
	%r7 = load i64, i64* %temp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr1, i32 0, i32 0), i64 %r7)
	br label %L8
L3:
	br label %L8
L8:
	; Exit (curr != 0)
	ret void
}

define i64 @main()
{
L9:
	%_mainretval = alloca i64
	; local variables
	; return
	store i64 0, i64* %_mainretval
	%r8 = load i64, i64* %_mainretval
	ret i64 %r8
	store i64 0, i64* %_mainretval
	%r9 = load i64, i64* %_mainretval
	ret i64 %r9
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [4 x i8] c"%ld\00", align 1