source_filename = "hard"
target triple = "x86_64-linux-gnu"

%struct.Cell = type {%struct.Cell*, i64}
@.nilCell = common global %struct.Cell* null
%struct.Row = type {%struct.Row*, %struct.Cell*}
@.nilRow = common global %struct.Row* null
%struct.ListNode = type {%struct.ListNode*, i64, i64}
@.nilListNode = common global %struct.ListNode* null

@matrix = common global %struct.Row* null
@maxrange = common global i64 0

define void @timesone(i64 %iter)
{
L0:
	%P_iter = alloca i64
	; local variables
	store i64 %iter, i64* %P_iter
	br label %L1
L1:
	; for (iter > 0)
	%r1 = load i64, i64* %P_iter
	%r2 = icmp sgt i64 %r1, 0
	br i1 %r2, label %L2, label %L3
L2:
	; maxrange = (maxrange * 1);
	%r3 = load i64, i64* @maxrange
	%r4 = mul i64 %r3, 1
	store i64 %r4, i64* @maxrange
	; iter = (iter - 1);
	%r5 = load i64, i64* %P_iter
	%r6 = sub i64 %r5, 1
	store i64 %r6, i64* %P_iter
	br label %L1
L3:
	ret void
}

define i64 @find(i64 %num, %struct.ListNode* %list)
{
L4:
	%find_retval = alloca i64
	%P_num = alloca i64
	; parameters
	%P_list = alloca %struct.ListNode*
	; local variables
	store i64 %num, i64* %P_num
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	br label %L5
L5:
	; If (list != nil)
	%r7 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r8 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r9 = icmp ne %struct.ListNode* %r7, %r8
	br i1 %r9, label %L6, label %L7
L6:
	; True (list != nil)
	br label %L8
L8:
	; If (list.lookup == num)
	%r10 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r11 = getelementptr %struct.ListNode, %struct.ListNode* %r10, i32 0, i32 1
	%r12 = load i64, i64* %r11
	%r13 = load i64, i64* %P_num
	%r14 = icmp eq i64 %r12, %r13
	br i1 %r14, label %L9, label %L10
L9:
	; True (list.lookup == num)
	; return list.value
	%r15 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r16 = getelementptr %struct.ListNode, %struct.ListNode* %r15, i32 0, i32 2
	%r17 = load i64, i64* %r16
	store i64 %r17, i64* %find_retval
	%r18 = load i64, i64* %find_retval
	ret i64 %r18
L10:
	; Else (list.lookup == num)
	; return find(num, list.next)
	%r19 = load i64, i64* %P_num
	%r20 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r21 = getelementptr %struct.ListNode, %struct.ListNode* %r20, i32 0, i32 0
	%r22 = load %struct.ListNode*, %struct.ListNode** %r21
	%r23 = call i64 @find(i64 %r19, %struct.ListNode* %r22)
	store i64 %r23, i64* %find_retval
	%r24 = load i64, i64* %find_retval
	ret i64 %r24
L7:
	br label %L11
L11:
	; Exit (list != nil)
	; return -1
	%r25 = sub i64 0, 1
	store i64 %r25, i64* %find_retval
	%r26 = load i64, i64* %find_retval
	ret i64 %r26
}

define %struct.ListNode* @add(i64 %lookup, i64 %value, %struct.ListNode* %list)
{
L12:
	%add_retval = alloca %struct.ListNode*
	%P_lookup = alloca i64
	; parameters
	%P_value = alloca i64
	%P_list = alloca %struct.ListNode*
	; local variables
	store i64 %lookup, i64* %P_lookup
	store i64 %value, i64* %P_value
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	br label %L13
L13:
	; If (list == nil)
	%r27 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r28 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r29 = icmp eq %struct.ListNode* %r27, %r28
	br i1 %r29, label %L14, label %L15
L14:
	; True (list == nil)
	; list = new ListNode;
	%r30 = call i8* @malloc(i32 24)
	%r31 = bitcast i8* %r30 to %struct.ListNode*
	store %struct.ListNode* %r31, %struct.ListNode** %P_list
	; list.lookup = lookup;
	%r32 = load i64, i64* %P_lookup
	%r33 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r34 = getelementptr %struct.ListNode, %struct.ListNode* %r33, i32 0, i32 1
	store i64 %r32, i64* %r34
	; list.value = value;
	%r35 = load i64, i64* %P_value
	%r36 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r37 = getelementptr %struct.ListNode, %struct.ListNode* %r36, i32 0, i32 2
	store i64 %r35, i64* %r37
	; list.next = nil;
	%r38 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r39 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r40 = getelementptr %struct.ListNode, %struct.ListNode* %r39, i32 0, i32 0
	store %struct.ListNode* %r38, %struct.ListNode** %r40
	br label %L16
L15:
	; Else (list == nil)
	; list.next = add(lookup, value, list.next);
	%r41 = load i64, i64* %P_lookup
	%r42 = load i64, i64* %P_value
	%r43 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r44 = getelementptr %struct.ListNode, %struct.ListNode* %r43, i32 0, i32 0
	%r45 = load %struct.ListNode*, %struct.ListNode** %r44
	%r46 = call %struct.ListNode* @add(i64 %r41, i64 %r42, %struct.ListNode* %r45)
	%r47 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r48 = getelementptr %struct.ListNode, %struct.ListNode* %r47, i32 0, i32 0
	store %struct.ListNode* %r46, %struct.ListNode** %r48
	br label %L16
L16:
	; Exit (list == nil)
	; return list
	%r49 = load %struct.ListNode*, %struct.ListNode** %P_list
	store %struct.ListNode* %r49, %struct.ListNode** %add_retval
	%r50 = load %struct.ListNode*, %struct.ListNode** %add_retval
	ret %struct.ListNode* %r50
}

define i64 @factorial(i64 %num, %struct.ListNode* %list)
{
L17:
	%factorial_retval = alloca i64
	%P_num = alloca i64
	; parameters
	%P_list = alloca %struct.ListNode*
	; local variables
	%lookup = alloca i64
	%fact = alloca i64
	store i64 %num, i64* %P_num
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	call void @timesone(i64 100)
	br label %L18
L18:
	; If (num == 1)
	%r51 = load i64, i64* %P_num
	%r52 = icmp eq i64 %r51, 1
	br i1 %r52, label %L19, label %L20
L19:
	; True (num == 1)
	; return 1
	store i64 1, i64* %factorial_retval
	%r53 = load i64, i64* %factorial_retval
	ret i64 %r53
L20:
	; Else (num == 1)
	; lookup = find(num, list);
	%r54 = load i64, i64* %P_num
	%r55 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r56 = call i64 @find(i64 %r54, %struct.ListNode* %r55)
	store i64 %r56, i64* %lookup
	br label %L21
L21:
	; If (lookup != -1)
	%r57 = load i64, i64* %lookup
	%r58 = sub i64 0, 1
	%r59 = icmp ne i64 %r57, %r58
	br i1 %r59, label %L22, label %L23
L22:
	; True (lookup != -1)
	; return lookup
	%r60 = load i64, i64* %lookup
	store i64 %r60, i64* %factorial_retval
	%r61 = load i64, i64* %factorial_retval
	ret i64 %r61
L23:
	br label %L24
L24:
	; Exit (lookup != -1)
	; fact = (num * factorial((num - 1), list));
	%r62 = load i64, i64* %P_num
	%r63 = load i64, i64* %P_num
	%r64 = sub i64 %r63, 1
	%r65 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r66 = call i64 @factorial(i64 %r64, %struct.ListNode* %r65)
	%r67 = mul i64 %r62, %r66
	store i64 %r67, i64* %fact
	br label %L25
L25:
	; If ((fact / 3) == 0)
	%r68 = load i64, i64* %fact
	%r69 = sdiv i64 %r68, 3
	%r70 = icmp eq i64 %r69, 0
	br i1 %r70, label %L26, label %L27
L26:
	; True ((fact / 3) == 0)
	%r71 = load i64, i64* %P_num
	%r72 = load i64, i64* %fact
	%r73 = load %struct.ListNode*, %struct.ListNode** %P_list
	call %struct.ListNode* @add(i64 %r71, i64 %r72, %struct.ListNode* %r73)
	br label %L28
L27:
	br label %L28
L28:
	; Exit ((fact / 3) == 0)
	; return fact
	%r74 = load i64, i64* %fact
	store i64 %r74, i64* %factorial_retval
	%r75 = load i64, i64* %factorial_retval
	ret i64 %r75
}

define void @maxfactorial(%struct.ListNode* %list, i64 %max)
{
L29:
	%P_list = alloca %struct.ListNode*
	; parameters
	%P_max = alloca i64
	; local variables
	%row = alloca %struct.Row*
	%cell = alloca %struct.Cell*
	%fact = alloca i64
	store %struct.ListNode* %list, %struct.ListNode** %P_list
	store i64 %max, i64* %P_max
	; list.next = nil;
	%r76 = load %struct.ListNode*, %struct.ListNode** @.nilListNode
	%r77 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r78 = getelementptr %struct.ListNode, %struct.ListNode* %r77, i32 0, i32 0
	store %struct.ListNode* %r76, %struct.ListNode** %r78
	; row = matrix;
	%r79 = load %struct.Row*, %struct.Row** @matrix
	store %struct.Row* %r79, %struct.Row** %row
	br label %L30
L30:
	; for (row != nil)
	%r80 = load %struct.Row*, %struct.Row** %row
	%r81 = load %struct.Row*, %struct.Row** @.nilRow
	%r82 = icmp ne %struct.Row* %r80, %r81
	br i1 %r82, label %L31, label %L39
L31:
	; cell = row.first;
	%r83 = load %struct.Row*, %struct.Row** %row
	%r84 = getelementptr %struct.Row, %struct.Row* %r83, i32 0, i32 1
	%r85 = load %struct.Cell*, %struct.Cell** %r84
	store %struct.Cell* %r85, %struct.Cell** %cell
	; row = row.next;
	%r86 = load %struct.Row*, %struct.Row** %row
	%r87 = getelementptr %struct.Row, %struct.Row* %r86, i32 0, i32 0
	%r88 = load %struct.Row*, %struct.Row** %r87
	store %struct.Row* %r88, %struct.Row** %row
	br label %L32
L32:
	; for (cell != nil)
	%r89 = load %struct.Cell*, %struct.Cell** %cell
	%r90 = load %struct.Cell*, %struct.Cell** @.nilCell
	%r91 = icmp ne %struct.Cell* %r89, %r90
	br i1 %r91, label %L33, label %L38
L33:
	; fact = factorial(cell.value, list);
	%r92 = load %struct.Cell*, %struct.Cell** %cell
	%r93 = getelementptr %struct.Cell, %struct.Cell* %r92, i32 0, i32 1
	%r94 = load i64, i64* %r93
	%r95 = load %struct.ListNode*, %struct.ListNode** %P_list
	%r96 = call i64 @factorial(i64 %r94, %struct.ListNode* %r95)
	store i64 %r96, i64* %fact
	; cell = cell.next;
	%r97 = load %struct.Cell*, %struct.Cell** %cell
	%r98 = getelementptr %struct.Cell, %struct.Cell* %r97, i32 0, i32 0
	%r99 = load %struct.Cell*, %struct.Cell** %r98
	store %struct.Cell* %r99, %struct.Cell** %cell
	br label %L34
L34:
	; If (fact > max)
	%r100 = load i64, i64* %fact
	%r101 = load i64, i64* %P_max
	%r102 = icmp sgt i64 %r100, %r101
	br i1 %r102, label %L35, label %L36
L35:
	; True (fact > max)
	; max = fact;
	%r103 = load i64, i64* %fact
	store i64 %r103, i64* %P_max
	br label %L37
L36:
	br label %L37
L37:
	; Exit (fact > max)
	br label %L32
L38:
	br label %L30
L39:
	%r104 = load i64, i64* %P_max
	call i32 (i8*, ...) @printf(i8* getelementptr inbounds ([18 x i8], [18 x i8]* @.fstr1, i32 0, i32 0), i64 %r104)
	ret void
}

define i64 @newvalue(i64 %height, i64 %width)
{
L40:
	%newvalue_retval = alloca i64
	%P_height = alloca i64
	; parameters
	%P_width = alloca i64
	; local variables
	%position = alloca i64
	store i64 %height, i64* %P_height
	store i64 %width, i64* %P_width
	call void @timesone(i64 0)
	; position = (height * width);
	%r105 = load i64, i64* %P_height
	%r106 = load i64, i64* %P_width
	%r107 = mul i64 %r105, %r106
	store i64 %r107, i64* %position
	; return ((maxrange / position) + height)
	%r108 = load i64, i64* @maxrange
	%r109 = load i64, i64* %position
	%r110 = sdiv i64 %r108, %r109
	%r111 = load i64, i64* %P_height
	%r112 = add i64 %r110, %r111
	store i64 %r112, i64* %newvalue_retval
	%r113 = load i64, i64* %newvalue_retval
	ret i64 %r113
}

define %struct.Cell* @newcell(%struct.Cell* %cell, i64 %height, i64 %width)
{
L41:
	%newcell_retval = alloca %struct.Cell*
	%P_cell = alloca %struct.Cell*
	; parameters
	%P_height = alloca i64
	%P_width = alloca i64
	; local variables
	store %struct.Cell* %cell, %struct.Cell** %P_cell
	store i64 %height, i64* %P_height
	store i64 %width, i64* %P_width
	; cell.value = newvalue(height, width);
	%r114 = load i64, i64* %P_height
	%r115 = load i64, i64* %P_width
	%r116 = call i64 @newvalue(i64 %r114, i64 %r115)
	%r117 = load %struct.Cell*, %struct.Cell** %P_cell
	%r118 = getelementptr %struct.Cell, %struct.Cell* %r117, i32 0, i32 1
	store i64 %r116, i64* %r118
	br label %L42
L42:
	; If (width > 1)
	%r119 = load i64, i64* %P_width
	%r120 = icmp sgt i64 %r119, 1
	br i1 %r120, label %L43, label %L44
L43:
	; True (width > 1)
	; cell.next = newcell(new Cell, height, (width - 1));
	%r121 = call i8* @malloc(i32 16)
	%r122 = bitcast i8* %r121 to %struct.Cell*
	%r123 = load i64, i64* %P_height
	%r124 = load i64, i64* %P_width
	%r125 = sub i64 %r124, 1
	%r126 = call %struct.Cell* @newcell(%struct.Cell* %r122, i64 %r123, i64 %r125)
	%r127 = load %struct.Cell*, %struct.Cell** %P_cell
	%r128 = getelementptr %struct.Cell, %struct.Cell* %r127, i32 0, i32 0
	store %struct.Cell* %r126, %struct.Cell** %r128
	br label %L45
L44:
	; Else (width > 1)
	; cell.next = nil;
	%r129 = load %struct.Cell*, %struct.Cell** @.nilCell
	%r130 = load %struct.Cell*, %struct.Cell** %P_cell
	%r131 = getelementptr %struct.Cell, %struct.Cell* %r130, i32 0, i32 0
	store %struct.Cell* %r129, %struct.Cell** %r131
	br label %L45
L45:
	; Exit (width > 1)
	; return cell
	%r132 = load %struct.Cell*, %struct.Cell** %P_cell
	store %struct.Cell* %r132, %struct.Cell** %newcell_retval
	%r133 = load %struct.Cell*, %struct.Cell** %newcell_retval
	ret %struct.Cell* %r133
}

define %struct.Row* @newrow(%struct.Row* %row, i64 %height, i64 %width)
{
L46:
	%newrow_retval = alloca %struct.Row*
	%P_row = alloca %struct.Row*
	; parameters
	%P_height = alloca i64
	%P_width = alloca i64
	; local variables
	store %struct.Row* %row, %struct.Row** %P_row
	store i64 %height, i64* %P_height
	store i64 %width, i64* %P_width
	; row.first = (newcell(new Cell, height, width));
	%r134 = call i8* @malloc(i32 16)
	%r135 = bitcast i8* %r134 to %struct.Cell*
	%r136 = load i64, i64* %P_height
	%r137 = load i64, i64* %P_width
	%r138 = call %struct.Cell* @newcell(%struct.Cell* %r135, i64 %r136, i64 %r137)
	%r139 = load %struct.Row*, %struct.Row** %P_row
	%r140 = getelementptr %struct.Row, %struct.Row* %r139, i32 0, i32 1
	store %struct.Cell* %r138, %struct.Cell** %r140
	br label %L47
L47:
	; If (height > 1)
	%r141 = load i64, i64* %P_height
	%r142 = icmp sgt i64 %r141, 1
	br i1 %r142, label %L48, label %L49
L48:
	; True (height > 1)
	; row.next = newrow(new Row, (height - 1), width);
	%r143 = call i8* @malloc(i32 16)
	%r144 = bitcast i8* %r143 to %struct.Row*
	%r145 = load i64, i64* %P_height
	%r146 = sub i64 %r145, 1
	%r147 = load i64, i64* %P_width
	%r148 = call %struct.Row* @newrow(%struct.Row* %r144, i64 %r146, i64 %r147)
	%r149 = load %struct.Row*, %struct.Row** %P_row
	%r150 = getelementptr %struct.Row, %struct.Row* %r149, i32 0, i32 0
	store %struct.Row* %r148, %struct.Row** %r150
	br label %L50
L49:
	; Else (height > 1)
	; row.next = nil;
	%r151 = load %struct.Row*, %struct.Row** @.nilRow
	%r152 = load %struct.Row*, %struct.Row** %P_row
	%r153 = getelementptr %struct.Row, %struct.Row* %r152, i32 0, i32 0
	store %struct.Row* %r151, %struct.Row** %r153
	br label %L50
L50:
	; Exit (height > 1)
	; return row
	%r154 = load %struct.Row*, %struct.Row** %P_row
	store %struct.Row* %r154, %struct.Row** %newrow_retval
	%r155 = load %struct.Row*, %struct.Row** %newrow_retval
	ret %struct.Row* %r155
}

define void @newmatrix(i64 %height, i64 %width)
{
L51:
	%P_height = alloca i64
	; parameters
	%P_width = alloca i64
	; local variables
	store i64 %height, i64* %P_height
	store i64 %width, i64* %P_width
	; matrix = newrow(new Row, height, width);
	%r156 = call i8* @malloc(i32 16)
	%r157 = bitcast i8* %r156 to %struct.Row*
	%r158 = load i64, i64* %P_height
	%r159 = load i64, i64* %P_width
	%r160 = call %struct.Row* @newrow(%struct.Row* %r157, i64 %r158, i64 %r159)
	store %struct.Row* %r160, %struct.Row** @matrix
	ret void
}

define i64 @getmatrixsize(i64 %matrixsize)
{
L52:
	%getmatrixsize_retval = alloca i64
	%P_matrixsize = alloca i64
	; local variables
	store i64 %matrixsize, i64* %P_matrixsize
	br label %L53
L53:
	; If (matrixsize <= 0)
	%r161 = load i64, i64* %P_matrixsize
	%r162 = icmp sle i64 %r161, 0
	br i1 %r162, label %L54, label %L55
L54:
	; True (matrixsize <= 0)
	; matrixsize = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %P_matrixsize)
	%r163 = load i64, i64* %P_matrixsize
	call i64 @getmatrixsize(i64 %r163)
	br label %L56
L55:
	; Else (matrixsize <= 0)
	; return matrixsize
	%r164 = load i64, i64* %P_matrixsize
	store i64 %r164, i64* %getmatrixsize_retval
	%r165 = load i64, i64* %getmatrixsize_retval
	ret i64 %r165
L56:
	; Exit (matrixsize <= 0)
	; return matrixsize
	%r166 = load i64, i64* %P_matrixsize
	store i64 %r166, i64* %getmatrixsize_retval
	%r167 = load i64, i64* %getmatrixsize_retval
	ret i64 %r167
}

define i64 @getmaxrange(i64 %maxrange)
{
L57:
	%getmaxrange_retval = alloca i64
	%P_maxrange = alloca i64
	; local variables
	store i64 %maxrange, i64* %P_maxrange
	br label %L58
L58:
	; If (maxrange <= 1)
	%r168 = load i64, i64* %P_maxrange
	%r169 = icmp sle i64 %r168, 1
	br i1 %r169, label %L59, label %L60
L59:
	; True (maxrange <= 1)
	; maxrange = scan;
	call i32 (i8*, ...) @scanf(i8* getelementptr inbounds ([4 x i8], [4 x i8]* @.read, i32 0, i32 0), i64* %P_maxrange)
	%r170 = load i64, i64* %P_maxrange
	call i64 @getmaxrange(i64 %r170)
	br label %L61
L60:
	; Else (maxrange <= 1)
	; return maxrange
	%r171 = load i64, i64* %P_maxrange
	store i64 %r171, i64* %getmaxrange_retval
	%r172 = load i64, i64* %getmaxrange_retval
	ret i64 %r172
L61:
	; Exit (maxrange <= 1)
	; return maxrange
	%r173 = load i64, i64* %P_maxrange
	store i64 %r173, i64* %getmaxrange_retval
	%r174 = load i64, i64* %getmaxrange_retval
	ret i64 %r174
}

define i64 @main()
{
L62:
	%_mainretval = alloca i64
	; local variables
	%height = alloca i64
	%width = alloca i64
	; height = 0;
	store i64 0, i64* %height
	; width = 0;
	store i64 0, i64* %width
	; maxrange = 0;
	store i64 0, i64* @maxrange
	; height = getmatrixsize(height);
	%r175 = load i64, i64* %height
	%r176 = call i64 @getmatrixsize(i64 %r175)
	store i64 %r176, i64* %height
	; width = height;
	%r177 = load i64, i64* %height
	store i64 %r177, i64* %width
	; maxrange = getmaxrange(maxrange);
	%r178 = load i64, i64* @maxrange
	%r179 = call i64 @getmaxrange(i64 %r178)
	store i64 %r179, i64* @maxrange
	%r180 = load i64, i64* %height
	%r181 = load i64, i64* %width
	call void @newmatrix(i64 %r180, i64 %r181)
	%r182 = call i8* @malloc(i32 24)
	%r183 = bitcast i8* %r182 to %struct.ListNode*
	call void @maxfactorial(%struct.ListNode* %r183, i64 0)
	store i64 0, i64* %_mainretval
	%r184 = load i64, i64* %_mainretval
	ret i64 %r184
}

declare i8* @malloc(i32)
declare void @free(i8*)
declare i32 @printf(i8*, ...)
declare i32 @scanf(i8*, ...)

@.fstr1 = private unnamed_addr constant [18 x i8] c"Max Factorial=%ld\00", align 1
@.read = private unnamed_addr constant [4 x i8] c"%ld\00", align 1