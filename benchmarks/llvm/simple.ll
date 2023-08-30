source_filename = "simple"
target triple = "abc"

%struct.Point2D = type {i64, i64}
@.nilPoint2D = common global %struct.Point2D* null

@globalInit = common global i64 0

define %struct.Point2D* @Init(i64 %initVal)
{
L0:
	%_Initretval = alloca %struct.Point2D*
	%P_initVal = alloca i64
	%newPt = alloca %struct.Point2D*
	store i64 %initVal, i64* %P_initVal
	%r1 = load %struct.Point2D*, %struct.Point2D** @.nilPoint2D
	store %struct.Point2D* %r1, %struct.Point2D** %newPt
	br label %L1
L1:
	%r2 = load i64, i64* %P_initVal
	%r3 = icmp sgt i64 %r2, 0
	br i1 %r3, label %L2, label %L3
L2:
	%r4 = call i8* @malloc(i32 16)
	%r5 = bitcast i8* %r4 to %struct.Point2D*
	store %struct.Point2D* %r5, %struct.Point2D** %newPt
	%r6 = load i64, i64* %P_initVal
	%r7 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r8 = getelementptr %struct.Point2D, %struct.Point2D* %r7, i32 0, i32 0
	store i64 %r6, i64* %r8
	%r9 = load i64, i64* %P_initVal
	%r10 = load %struct.Point2D*, %struct.Point2D** %newPt
	%r11 = getelementptr %struct.Point2D, %struct.Point2D* %r10, i32 0, i32 1
	store i64 %r9, i64* %r11
	%r12 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r12, %struct.Point2D** %_Initretval
	%r13 = load %struct.Point2D*, %struct.Point2D** %_Initretval
	ret %struct.Point2D* %r13
L3:
	br label %L4
L4:
	%r14 = load %struct.Point2D*, %struct.Point2D** %newPt
	store %struct.Point2D* %r14, %struct.Point2D** %_Initretval
	%r15 = load %struct.Point2D*, %struct.Point2D** %_Initretval
	ret %struct.Point2D* %r15
}

define i64 @main()
{
L5:
	%_mainretval = alloca i64
	%a = alloca i64
	%b = alloca i64
	%pt1 = alloca %struct.Point2D*
	%pt2 = alloca %struct.Point2D*
	store i64 5, i64* %a
	%r16 = load i64, i64* %a
	%r17 = add i64 %r16, 7
	%r18 = mul i64 %r17, 3
	store i64 %r18, i64* %b
	%r19 = call i8* @malloc(i32 16)
	%r20 = bitcast i8* %r19 to %struct.Point2D*
	store %struct.Point2D* %r20, %struct.Point2D** %pt1
	%r21 = load i64, i64* %a
	%r22 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r23 = getelementptr %struct.Point2D, %struct.Point2D* %r22, i32 0, i32 0
	store i64 %r21, i64* %r23
	%r24 = load i64, i64* %b
	%r25 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r26 = getelementptr %struct.Point2D, %struct.Point2D* %r25, i32 0, i32 1
	store i64 %r24, i64* %r26
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* @globalInit)
	%r27 = load i64, i64* @globalInit
	%r28 = call %struct.Point2D* @Init(i64 %r27)
	store %struct.Point2D* %r28, %struct.Point2D** %pt2
	%r29 = load i64, i64* @globalInit
	%r30 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r31 = getelementptr %struct.Point2D, %struct.Point2D* %r30, i32 0, i32 0
	%r32 = load i64, i64* %r31
	%r33 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r34 = getelementptr %struct.Point2D, %struct.Point2D* %r33, i32 0, i32 1
	%r35 = load i64, i64* %r34
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([32 x i8], [32 x i8]* @.fstr1, i32 0, i32 0), i64 %r29, i64 %r32, i64 %r35)
	%r36 = load %struct.Point2D*, %struct.Point2D** %pt1
	%r37 = bitcast %struct.Point2D* %r36 to i8*
	call void @free(i8* %r37)
	%r38 = load %struct.Point2D*, %struct.Point2D** %pt2
	%r39 = bitcast %struct.Point2D* %r38 to i8*
	call void @free(i8* %r39)
	store i64 0, i64* %_mainretval
	%r40 = load i64, i64* %_mainretval
	ret i64 %r40
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [32 x i8] c"offset=%ld\0Apt2.x=%ld\0Apt2.y=%ld\0A\00", align 1