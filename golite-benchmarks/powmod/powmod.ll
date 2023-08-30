source_filename = "powmod"
target triple = "x86_64-linux-gnu"

%struct.MESSAGE = type {i64, %struct.MESSAGE*}
@.nilMESSAGE = common global %struct.MESSAGE* null


define i64 @mod(i64 %val, i64 %t)
{
L0:
	%mod_retval = alloca i64
	%P_val = alloca i64
	; parameters
	%P_t = alloca i64
	; local variables
	%temp = alloca i64
	store i64 %val, i64* %P_val
	store i64 %t, i64* %P_t
	; temp = (val / t);
	%r1 = load i64, i64* %P_val
	%r2 = load i64, i64* %P_t
	%r3 = sdiv i64 %r1, %r2
	store i64 %r3, i64* %temp
	; return (val - (temp * t))
	%r4 = load i64, i64* %P_val
	%r5 = load i64, i64* %temp
	%r6 = load i64, i64* %P_t
	%r7 = mul i64 %r5, %r6
	%r8 = sub i64 %r4, %r7
	store i64 %r8, i64* %mod_retval
	%r9 = load i64, i64* %mod_retval
	ret i64 %r9
}

define i64 @power(i64 %x, i64 %n)
{
L1:
	%power_retval = alloca i64
	%P_x = alloca i64
	; parameters
	%P_n = alloca i64
	; local variables
	%temp = alloca i64
	store i64 %x, i64* %P_x
	store i64 %n, i64* %P_n
	br label %L2
L2:
	; If (n == 0)
	%r10 = load i64, i64* %P_n
	%r11 = icmp eq i64 %r10, 0
	br i1 %r11, label %L3, label %L4
L3:
	; True (n == 0)
	; return 1
	store i64 1, i64* %power_retval
	%r12 = load i64, i64* %power_retval
	ret i64 %r12
L4:
	; Else (n == 0)
	br label %L5
L5:
	; If (mod(n, 2) == 1)
	%r13 = load i64, i64* %P_n
	%r14 = call i64 @mod(i64 %r13, i64 2)
	%r15 = icmp eq i64 %r14, 1
	br i1 %r15, label %L6, label %L7
L6:
	; True (mod(n, 2) == 1)
	; return (x * power(x, (n - 1)))
	%r16 = load i64, i64* %P_x
	%r17 = load i64, i64* %P_x
	%r18 = load i64, i64* %P_n
	%r19 = sub i64 %r18, 1
	%r20 = call i64 @power(i64 %r17, i64 %r19)
	%r21 = mul i64 %r16, %r20
	store i64 %r21, i64* %power_retval
	%r22 = load i64, i64* %power_retval
	ret i64 %r22
L7:
	; Else (mod(n, 2) == 1)
	; temp = power(x, (n / 2));
	%r23 = load i64, i64* %P_x
	%r24 = load i64, i64* %P_n
	%r25 = sdiv i64 %r24, 2
	%r26 = call i64 @power(i64 %r23, i64 %r25)
	store i64 %r26, i64* %temp
	; return (temp * temp)
	%r27 = load i64, i64* %temp
	%r28 = load i64, i64* %temp
	%r29 = mul i64 %r27, %r28
	store i64 %r29, i64* %power_retval
	%r30 = load i64, i64* %power_retval
	ret i64 %r30
}

define void @crypt(i64 %m, i64 %key, %struct.MESSAGE* %msg)
{
L8:
	%P_m = alloca i64
	; parameters
	%P_key = alloca i64
	%P_msg = alloca %struct.MESSAGE*
	; local variables
	store i64 %m, i64* %P_m
	store i64 %key, i64* %P_key
	store %struct.MESSAGE* %msg, %struct.MESSAGE** %P_msg
	; msg.val = mod(power(msg.val, key), m);
	%r31 = load %struct.MESSAGE*, %struct.MESSAGE** %P_msg
	%r32 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r31, i32 0, i32 0
	%r33 = load i64, i64* %r32
	%r34 = load i64, i64* %P_key
	%r35 = call i64 @power(i64 %r33, i64 %r34)
	%r36 = load i64, i64* %P_m
	%r37 = call i64 @mod(i64 %r35, i64 %r36)
	%r38 = load %struct.MESSAGE*, %struct.MESSAGE** %P_msg
	%r39 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r38, i32 0, i32 0
	store i64 %r37, i64* %r39
	br label %L9
L9:
	; If (msg.next != nil)
	%r40 = load %struct.MESSAGE*, %struct.MESSAGE** %P_msg
	%r41 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r40, i32 0, i32 1
	%r42 = load %struct.MESSAGE*, %struct.MESSAGE** %r41
	%r43 = load %struct.MESSAGE*, %struct.MESSAGE** @.nilMESSAGE
	%r44 = icmp ne %struct.MESSAGE* %r42, %r43
	br i1 %r44, label %L10, label %L11
L10:
	; True (msg.next != nil)
	%r45 = load i64, i64* %P_m
	%r46 = load i64, i64* %P_key
	%r47 = load %struct.MESSAGE*, %struct.MESSAGE** %P_msg
	%r48 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r47, i32 0, i32 1
	%r49 = load %struct.MESSAGE*, %struct.MESSAGE** %r48
	call void @crypt(i64 %r45, i64 %r46, %struct.MESSAGE* %r49)
	br label %L12
L11:
	br label %L12
L12:
	; Exit (msg.next != nil)
	ret void
}

define i64 @main()
{
L13:
	%_mainretval = alloca i64
	; local variables
	%current = alloca %struct.MESSAGE*
	%temp = alloca %struct.MESSAGE*
	%key = alloca i64
	%mod = alloca i64
	%length = alloca i64
	%readTemp = alloca i64
	%printTemp = alloca i64
	%start = alloca %struct.MESSAGE*
	; start = new MESSAGE;
	%r50 = call i8* @malloc(i32 16)
	%r51 = bitcast i8* %r50 to %struct.MESSAGE*
	store %struct.MESSAGE* %r51, %struct.MESSAGE** %start
	; current = start;
	%r52 = load %struct.MESSAGE*, %struct.MESSAGE** %start
	store %struct.MESSAGE* %r52, %struct.MESSAGE** %current
	; key = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %key)
	; mod = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %mod)
	; length = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %length)
	; length = (length - 1);
	%r53 = load i64, i64* %length
	%r54 = sub i64 %r53, 1
	store i64 %r54, i64* %length
	; readTemp = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %readTemp)
	; current.val = readTemp;
	%r55 = load i64, i64* %readTemp
	%r56 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r57 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r56, i32 0, i32 0
	store i64 %r55, i64* %r57
	br label %L14
L14:
	; for (length > 0)
	%r58 = load i64, i64* %length
	%r59 = icmp sgt i64 %r58, 0
	br i1 %r59, label %L15, label %L16
L15:
	; current.next = new MESSAGE;
	%r60 = call i8* @malloc(i32 16)
	%r61 = bitcast i8* %r60 to %struct.MESSAGE*
	%r62 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r63 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r62, i32 0, i32 1
	store %struct.MESSAGE* %r61, %struct.MESSAGE** %r63
	; current = current.next;
	%r64 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r65 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r64, i32 0, i32 1
	%r66 = load %struct.MESSAGE*, %struct.MESSAGE** %r65
	store %struct.MESSAGE* %r66, %struct.MESSAGE** %current
	; readTemp = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %readTemp)
	; current.val = readTemp;
	%r67 = load i64, i64* %readTemp
	%r68 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r69 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r68, i32 0, i32 0
	store i64 %r67, i64* %r69
	; length = (length - 1);
	%r70 = load i64, i64* %length
	%r71 = sub i64 %r70, 1
	store i64 %r71, i64* %length
	br label %L14
L16:
	; current.next = nil;
	%r72 = load %struct.MESSAGE*, %struct.MESSAGE** @.nilMESSAGE
	%r73 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r74 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r73, i32 0, i32 1
	store %struct.MESSAGE* %r72, %struct.MESSAGE** %r74
	%r75 = load i64, i64* %mod
	%r76 = load i64, i64* %key
	%r77 = load %struct.MESSAGE*, %struct.MESSAGE** %start
	call void @crypt(i64 %r75, i64 %r76, %struct.MESSAGE* %r77)
	; current = start;
	%r78 = load %struct.MESSAGE*, %struct.MESSAGE** %start
	store %struct.MESSAGE* %r78, %struct.MESSAGE** %current
	br label %L17
L17:
	; for (current != nil)
	%r79 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r80 = load %struct.MESSAGE*, %struct.MESSAGE** @.nilMESSAGE
	%r81 = icmp ne %struct.MESSAGE* %r79, %r80
	br i1 %r81, label %L18, label %L19
L18:
	; temp = current;
	%r82 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	store %struct.MESSAGE* %r82, %struct.MESSAGE** %temp
	; printTemp = current.val;
	%r83 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r84 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r83, i32 0, i32 0
	%r85 = load i64, i64* %r84
	store i64 %r85, i64* %printTemp
	%r86 = load i64, i64* %printTemp
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([5 x i8], [5 x i8]* @.fstr1, i32 0, i32 0), i64 %r86)
	; current = current.next;
	%r87 = load %struct.MESSAGE*, %struct.MESSAGE** %current
	%r88 = getelementptr %struct.MESSAGE, %struct.MESSAGE* %r87, i32 0, i32 1
	%r89 = load %struct.MESSAGE*, %struct.MESSAGE** %r88
	store %struct.MESSAGE* %r89, %struct.MESSAGE** %current
	%r90 = load %struct.MESSAGE*, %struct.MESSAGE** %temp
	%r91 = bitcast %struct.MESSAGE* %r90 to i8*
	call void @free(i8* %r91)
	br label %L17
L19:
	store i64 0, i64* %_mainretval
	%r92 = load i64, i64* %_mainretval
	ret i64 %r92
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1
@.fstr1 = private unnamed_addr constant [5 x i8] c"%ld\0A\00", align 1