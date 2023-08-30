	.text
	.file	"hard"
	.globl	timesone                # -- Begin function timesone
	.p2align	4, 0x90
	.type	timesone,@function
timesone:                               # @timesone
	.cfi_startproc
# %bb.0:                                # %L0
	movq	%rdi, -8(%rsp)
	cmpq	$0, -8(%rsp)
	jle	.LBB0_3
	.p2align	4, 0x90
.LBB0_2:                                # %L2
                                        # =>This Inner Loop Header: Depth=1
	decq	-8(%rsp)
	cmpq	$0, -8(%rsp)
	jg	.LBB0_2
.LBB0_3:                                # %L3
	retq
.Lfunc_end0:
	.size	timesone, .Lfunc_end0-timesone
	.cfi_endproc
                                        # -- End function
	.globl	find                    # -- Begin function find
	.p2align	4, 0x90
	.type	find,@function
find:                                   # @find
	.cfi_startproc
# %bb.0:                                # %L4
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	movq	%rsi, (%rsp)
	cmpq	.nilListNode(%rip), %rsi
	je	.LBB1_5
# %bb.1:                                # %L6
	movq	(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	8(%rsp), %rax
	jne	.LBB1_4
# %bb.2:                                # %L9
	movq	(%rsp), %rax
	movq	16(%rax), %rax
	jmp	.LBB1_3
.LBB1_5:                                # %L7
	movq	$-1, 16(%rsp)
	movq	$-1, %rax
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_4:                                # %L10
	.cfi_def_cfa_offset 32
	movq	8(%rsp), %rdi
	movq	(%rsp), %rax
	movq	(%rax), %rsi
	callq	find
.LBB1_3:                                # %L9
	movq	%rax, 16(%rsp)
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	find, .Lfunc_end1-find
	.cfi_endproc
                                        # -- End function
	.globl	add                     # -- Begin function add
	.p2align	4, 0x90
	.type	add,@function
add:                                    # @add
	.cfi_startproc
# %bb.0:                                # %L12
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 24(%rsp)
	movq	%rsi, 16(%rsp)
	movq	%rdx, 8(%rsp)
	cmpq	.nilListNode(%rip), %rdx
	je	.LBB2_1
# %bb.2:                                # %L15
	movq	24(%rsp), %rdi
	movq	16(%rsp), %rsi
	movq	8(%rsp), %rax
	movq	(%rax), %rdx
	callq	add
	jmp	.LBB2_3
.LBB2_1:                                # %L14
	movl	$24, %edi
	callq	malloc
	movq	%rax, 8(%rsp)
	movq	24(%rsp), %rcx
	movq	%rcx, 8(%rax)
	movq	16(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 16(%rcx)
	movq	.nilListNode(%rip), %rax
.LBB2_3:                                # %L16
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	movq	8(%rsp), %rax
	movq	%rax, 32(%rsp)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	add, .Lfunc_end2-add
	.cfi_endproc
                                        # -- End function
	.globl	factorial               # -- Begin function factorial
	.p2align	4, 0x90
	.type	factorial,@function
factorial:                              # @factorial
	.cfi_startproc
# %bb.0:                                # %L17
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$48, %rsp
	.cfi_def_cfa_offset 64
	.cfi_offset %rbx, -16
	movq	%rdi, 8(%rsp)
	movq	%rsi, 16(%rsp)
	movl	$100, %edi
	callq	timesone
	cmpq	$1, 8(%rsp)
	jne	.LBB3_2
# %bb.1:                                # %L19
	movq	$1, 32(%rsp)
	movl	$1, %eax
	jmp	.LBB3_5
.LBB3_2:                                # %L20
	movq	8(%rsp), %rdi
	movq	16(%rsp), %rsi
	callq	find
	movq	%rax, 40(%rsp)
	cmpq	$-1, %rax
	je	.LBB3_6
# %bb.3:                                # %L22
	movq	40(%rsp), %rax
	jmp	.LBB3_4
.LBB3_6:                                # %L23
	movq	8(%rsp), %rbx
	leaq	-1(%rbx), %rdi
	movq	16(%rsp), %rsi
	callq	factorial
	imulq	%rbx, %rax
	movq	%rax, 24(%rsp)
	movabsq	$6148914691236517206, %rcx # imm = 0x5555555555555556
	imulq	%rcx
	movq	%rdx, %rax
	shrq	$63, %rax
	addq	%rdx, %rax
	jne	.LBB3_8
# %bb.7:                                # %L26
	movq	8(%rsp), %rdi
	movq	24(%rsp), %rsi
	movq	16(%rsp), %rdx
	callq	add
.LBB3_8:                                # %L28
	movq	24(%rsp), %rax
.LBB3_4:                                # %L22
	movq	%rax, 32(%rsp)
.LBB3_5:                                # %L22
	addq	$48, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	factorial, .Lfunc_end3-factorial
	.cfi_endproc
                                        # -- End function
	.globl	maxfactorial            # -- Begin function maxfactorial
	.p2align	4, 0x90
	.type	maxfactorial,@function
maxfactorial:                           # @maxfactorial
	.cfi_startproc
# %bb.0:                                # %L29
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 32(%rsp)
	movq	%rsi, 16(%rsp)
	movq	.nilListNode(%rip), %rax
	movq	%rax, (%rdi)
	movq	matrix(%rip), %rax
	movq	%rax, 8(%rsp)
	.p2align	4, 0x90
.LBB4_1:                                # %L30
                                        # =>This Loop Header: Depth=1
                                        #     Child Loop BB4_3 Depth 2
	movq	8(%rsp), %rax
	cmpq	.nilRow(%rip), %rax
	je	.LBB4_6
# %bb.2:                                # %L31
                                        #   in Loop: Header=BB4_1 Depth=1
	movq	8(%rsp), %rax
	movq	8(%rax), %rcx
	movq	%rcx, (%rsp)
	movq	(%rax), %rax
	movq	%rax, 8(%rsp)
	.p2align	4, 0x90
.LBB4_3:                                # %L32
                                        #   Parent Loop BB4_1 Depth=1
                                        # =>  This Inner Loop Header: Depth=2
	movq	(%rsp), %rax
	cmpq	.nilCell(%rip), %rax
	je	.LBB4_1
# %bb.4:                                # %L33
                                        #   in Loop: Header=BB4_3 Depth=2
	movq	(%rsp), %rax
	movq	8(%rax), %rdi
	movq	32(%rsp), %rsi
	callq	factorial
	movq	%rax, 24(%rsp)
	movq	(%rsp), %rcx
	movq	(%rcx), %rcx
	movq	%rcx, (%rsp)
	cmpq	16(%rsp), %rax
	jle	.LBB4_3
# %bb.5:                                # %L35
                                        #   in Loop: Header=BB4_3 Depth=2
	movq	24(%rsp), %rax
	movq	%rax, 16(%rsp)
	jmp	.LBB4_3
.LBB4_6:                                # %L39
	movq	16(%rsp), %rsi
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end4:
	.size	maxfactorial, .Lfunc_end4-maxfactorial
	.cfi_endproc
                                        # -- End function
	.globl	newvalue                # -- Begin function newvalue
	.p2align	4, 0x90
	.type	newvalue,@function
newvalue:                               # @newvalue
	.cfi_startproc
# %bb.0:                                # %L40
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 16(%rsp)
	movq	%rsi, 8(%rsp)
	xorl	%edi, %edi
	callq	timesone
	movq	16(%rsp), %rcx
	movq	8(%rsp), %rsi
	imulq	%rcx, %rsi
	movq	%rsi, 24(%rsp)
	movq	maxrange(%rip), %rax
	cqto
	idivq	%rsi
	addq	%rcx, %rax
	movq	%rax, 32(%rsp)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end5:
	.size	newvalue, .Lfunc_end5-newvalue
	.cfi_endproc
                                        # -- End function
	.globl	newcell                 # -- Begin function newcell
	.p2align	4, 0x90
	.type	newcell,@function
newcell:                                # @newcell
	.cfi_startproc
# %bb.0:                                # %L41
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 8(%rsp)
	movq	%rsi, 24(%rsp)
	movq	%rdx, 16(%rsp)
	movq	%rsi, %rdi
	movq	%rdx, %rsi
	callq	newvalue
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	cmpq	$2, 16(%rsp)
	jl	.LBB6_2
# %bb.1:                                # %L43
	movl	$16, %edi
	callq	malloc
	movq	24(%rsp), %rsi
	movq	16(%rsp), %rdx
	decq	%rdx
	movq	%rax, %rdi
	callq	newcell
	jmp	.LBB6_3
.LBB6_2:                                # %L44
	movq	.nilCell(%rip), %rax
.LBB6_3:                                # %L45
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	movq	8(%rsp), %rax
	movq	%rax, 32(%rsp)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end6:
	.size	newcell, .Lfunc_end6-newcell
	.cfi_endproc
                                        # -- End function
	.globl	newrow                  # -- Begin function newrow
	.p2align	4, 0x90
	.type	newrow,@function
newrow:                                 # @newrow
	.cfi_startproc
# %bb.0:                                # %L46
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 8(%rsp)
	movq	%rsi, 16(%rsp)
	movq	%rdx, 24(%rsp)
	movl	$16, %edi
	callq	malloc
	movq	16(%rsp), %rsi
	movq	24(%rsp), %rdx
	movq	%rax, %rdi
	callq	newcell
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	cmpq	$2, 16(%rsp)
	jl	.LBB7_2
# %bb.1:                                # %L48
	movl	$16, %edi
	callq	malloc
	movq	16(%rsp), %rsi
	decq	%rsi
	movq	24(%rsp), %rdx
	movq	%rax, %rdi
	callq	newrow
	jmp	.LBB7_3
.LBB7_2:                                # %L49
	movq	.nilRow(%rip), %rax
.LBB7_3:                                # %L50
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	movq	8(%rsp), %rax
	movq	%rax, 32(%rsp)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end7:
	.size	newrow, .Lfunc_end7-newrow
	.cfi_endproc
                                        # -- End function
	.globl	newmatrix               # -- Begin function newmatrix
	.p2align	4, 0x90
	.type	newmatrix,@function
newmatrix:                              # @newmatrix
	.cfi_startproc
# %bb.0:                                # %L51
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 16(%rsp)
	movq	%rsi, 8(%rsp)
	movl	$16, %edi
	callq	malloc
	movq	16(%rsp), %rsi
	movq	8(%rsp), %rdx
	movq	%rax, %rdi
	callq	newrow
	movq	%rax, matrix(%rip)
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end8:
	.size	newmatrix, .Lfunc_end8-newmatrix
	.cfi_endproc
                                        # -- End function
	.globl	getmatrixsize           # -- Begin function getmatrixsize
	.p2align	4, 0x90
	.type	getmatrixsize,@function
getmatrixsize:                          # @getmatrixsize
	.cfi_startproc
# %bb.0:                                # %L52
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	cmpq	$0, 8(%rsp)
	jg	.LBB9_2
# %bb.1:                                # %L54
	leaq	8(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movq	8(%rsp), %rdi
	callq	getmatrixsize
.LBB9_2:                                # %L56
	movq	8(%rsp), %rax
	movq	%rax, 16(%rsp)
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end9:
	.size	getmatrixsize, .Lfunc_end9-getmatrixsize
	.cfi_endproc
                                        # -- End function
	.globl	getmaxrange             # -- Begin function getmaxrange
	.p2align	4, 0x90
	.type	getmaxrange,@function
getmaxrange:                            # @getmaxrange
	.cfi_startproc
# %bb.0:                                # %L57
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	cmpq	$1, 8(%rsp)
	jg	.LBB10_2
# %bb.1:                                # %L59
	leaq	8(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movq	8(%rsp), %rdi
	callq	getmaxrange
.LBB10_2:                               # %L61
	movq	8(%rsp), %rax
	movq	%rax, 16(%rsp)
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end10:
	.size	getmaxrange, .Lfunc_end10-getmaxrange
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L62
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	$0, 8(%rsp)
	movq	$0, (%rsp)
	movq	$0, maxrange(%rip)
	xorl	%edi, %edi
	callq	getmatrixsize
	movq	%rax, 8(%rsp)
	movq	%rax, (%rsp)
	movq	maxrange(%rip), %rdi
	callq	getmaxrange
	movq	%rax, maxrange(%rip)
	movq	8(%rsp), %rdi
	movq	(%rsp), %rsi
	callq	newmatrix
	movl	$24, %edi
	callq	malloc
	movq	%rax, %rdi
	xorl	%esi, %esi
	callq	maxfactorial
	movq	$0, 16(%rsp)
	xorl	%eax, %eax
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end11:
	.size	main, .Lfunc_end11-main
	.cfi_endproc
                                        # -- End function
	.type	.nilCell,@object        # @.nilCell
	.comm	.nilCell,8,8
	.type	.nilRow,@object         # @.nilRow
	.comm	.nilRow,8,8
	.type	.nilListNode,@object    # @.nilListNode
	.comm	.nilListNode,8,8
	.type	matrix,@object          # @matrix
	.comm	matrix,8,8
	.type	maxrange,@object        # @maxrange
	.comm	maxrange,8,8
	.type	.L.fstr1,@object        # @.fstr1
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.fstr1:
	.asciz	"Max Factorial=%ld"
	.size	.L.fstr1, 18

	.type	.L.read,@object         # @.read
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.section	".note.GNU-stack","",@progbits
