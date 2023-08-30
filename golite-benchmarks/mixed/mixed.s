	.text
	.file	"mixed"
	.globl	tailrecursive           # -- Begin function tailrecursive
	.p2align	4, 0x90
	.type	tailrecursive,@function
tailrecursive:                          # @tailrecursive
	.cfi_startproc
# %bb.0:                                # %L0
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	testq	%rdi, %rdi
	jg	.LBB0_2
# %bb.1:                                # %L2
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB0_2:                                # %L3
	.cfi_def_cfa_offset 32
	movl	$24, %edi
	callq	malloc
	movq	%rax, 16(%rsp)
	movq	%rax, unusedGlobal(%rip)
	movq	8(%rsp), %rdi
	decq	%rdi
	callq	tailrecursive
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	tailrecursive, .Lfunc_end0-tailrecursive
	.cfi_endproc
                                        # -- End function
	.globl	add                     # -- Begin function add
	.p2align	4, 0x90
	.type	add,@function
add:                                    # @add
	.cfi_startproc
# %bb.0:                                # %L5
	movq	%rdi, %rax
	movq	%rdi, -16(%rsp)
	movq	%rsi, -24(%rsp)
	addq	%rsi, %rax
	movq	%rax, -8(%rsp)
	retq
.Lfunc_end1:
	.size	add, .Lfunc_end1-add
	.cfi_endproc
                                        # -- End function
	.globl	domath                  # -- Begin function domath
	.p2align	4, 0x90
	.type	domath,@function
domath:                                 # @domath
	.cfi_startproc
# %bb.0:                                # %L6
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 32(%rsp)
	movl	$24, %edi
	callq	malloc
	movq	%rax, 16(%rsp)
	movl	$8, %edi
	callq	malloc
	movq	16(%rsp), %rcx
	movq	%rax, 16(%rcx)
	movl	$24, %edi
	callq	malloc
	movq	%rax, 8(%rsp)
	movl	$8, %edi
	callq	malloc
	movq	8(%rsp), %rcx
	movq	%rax, 16(%rcx)
	movq	32(%rsp), %rax
	movq	16(%rsp), %rcx
	movq	%rax, (%rcx)
	movq	8(%rsp), %rax
	movq	$3, (%rax)
	movq	16(%rsp), %rax
	movq	(%rax), %rcx
	movq	16(%rax), %rax
	movq	%rcx, (%rax)
	movq	8(%rsp), %rax
	movq	(%rax), %rcx
	movq	16(%rax), %rax
	movq	%rcx, (%rax)
	cmpq	$0, 32(%rsp)
	jle	.LBB2_3
	.p2align	4, 0x90
.LBB2_2:                                # %L8
                                        # =>This Inner Loop Header: Depth=1
	movq	16(%rsp), %rcx
	movq	(%rcx), %rax
	movq	8(%rsp), %rsi
	imulq	(%rsi), %rax
	movq	%rax, 24(%rsp)
	movq	16(%rcx), %rdx
	imulq	(%rdx), %rax
	cqto
	idivq	(%rsi)
	movq	%rax, 24(%rsp)
	movq	16(%rsi), %rax
	movq	(%rax), %rdi
	movq	(%rcx), %rsi
	callq	add
	movq	%rax, 24(%rsp)
	movq	8(%rsp), %rax
	movq	(%rax), %rax
	movq	16(%rsp), %rcx
	subq	(%rcx), %rax
	movq	%rax, 24(%rsp)
	decq	32(%rsp)
	cmpq	$0, 32(%rsp)
	jg	.LBB2_2
.LBB2_3:                                # %L9
	movq	16(%rsp), %rdi
	callq	free
	movq	8(%rsp), %rdi
	callq	free
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	domath, .Lfunc_end2-domath
	.cfi_endproc
                                        # -- End function
	.globl	objinstantiation        # -- Begin function objinstantiation
	.p2align	4, 0x90
	.type	objinstantiation,@function
objinstantiation:                       # @objinstantiation
	.cfi_startproc
# %bb.0:                                # %L10
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	cmpq	$0, 8(%rsp)
	jle	.LBB3_3
	.p2align	4, 0x90
.LBB3_2:                                # %L12
                                        # =>This Inner Loop Header: Depth=1
	movl	$24, %edi
	callq	malloc
	movq	%rax, 16(%rsp)
	movq	%rax, %rdi
	callq	free
	decq	8(%rsp)
	cmpq	$0, 8(%rsp)
	jg	.LBB3_2
.LBB3_3:                                # %L13
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	objinstantiation, .Lfunc_end3-objinstantiation
	.cfi_endproc
                                        # -- End function
	.globl	ackermann               # -- Begin function ackermann
	.p2align	4, 0x90
	.type	ackermann,@function
ackermann:                              # @ackermann
	.cfi_startproc
# %bb.0:                                # %L14
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$32, %rsp
	.cfi_def_cfa_offset 48
	.cfi_offset %rbx, -16
	movq	%rdi, 16(%rsp)
	movq	%rsi, 8(%rsp)
	testq	%rdi, %rdi
	je	.LBB4_3
# %bb.1:                                # %L17
	cmpq	$0, 8(%rsp)
	je	.LBB4_4
# %bb.2:                                # %L21
	movq	16(%rsp), %rdi
	movq	8(%rsp), %rsi
	decq	%rsi
	leaq	-1(%rdi), %rbx
	callq	ackermann
	movq	%rbx, %rdi
	movq	%rax, %rsi
	jmp	.LBB4_5
.LBB4_3:                                # %L16
	movq	8(%rsp), %rax
	incq	%rax
	jmp	.LBB4_6
.LBB4_4:                                # %L20
	movq	16(%rsp), %rdi
	decq	%rdi
	movl	$1, %esi
.LBB4_5:                                # %L16
	callq	ackermann
.LBB4_6:                                # %L16
	movq	%rax, 24(%rsp)
	addq	$32, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end4:
	.size	ackermann, .Lfunc_end4-ackermann
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L22
	subq	$56, %rsp
	.cfi_def_cfa_offset 64
	leaq	16(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	8(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movq	%rsp, %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	32(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	24(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movq	16(%rsp), %rdi
	callq	tailrecursive
	movq	8(%rsp), %rdi
	callq	domath
	movq	(%rsp), %rdi
	callq	objinstantiation
	movq	32(%rsp), %rdi
	movq	24(%rsp), %rsi
	callq	ackermann
	movq	%rax, 40(%rsp)
	movq	16(%rsp), %rsi
	movq	8(%rsp), %rdx
	movq	(%rsp), %rcx
	movl	$.L.fstr1, %edi
	movq	%rax, %r8
	xorl	%eax, %eax
	callq	printf
	movq	$0, 48(%rsp)
	xorl	%eax, %eax
	addq	$56, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end5:
	.size	main, .Lfunc_end5-main
	.cfi_endproc
                                        # -- End function
	.type	.nilsimple,@object      # @.nilsimple
	.comm	.nilsimple,8,8
	.type	.nilfoo,@object         # @.nilfoo
	.comm	.nilfoo,8,8
	.type	globalfoo,@object       # @globalfoo
	.comm	globalfoo,8,8
	.type	unusedGlobal,@object    # @unusedGlobal
	.comm	unusedGlobal,8,8
	.type	.L.read,@object         # @.read
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.type	.L.fstr1,@object        # @.fstr1
.L.fstr1:
	.asciz	"a=%ld\nb=%ld\nc=%ld,temp=%ld"
	.size	.L.fstr1, 27

	.section	".note.GNU-stack","",@progbits
