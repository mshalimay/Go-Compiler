	.text
	.file	"thermopylae"
	.globl	fact                    # -- Begin function fact
	.p2align	4, 0x90
	.type	fact,@function
fact:                                   # @fact
	.cfi_startproc
# %bb.0:                                # %L0
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$16, %rsp
	.cfi_def_cfa_offset 32
	.cfi_offset %rbx, -16
	movq	%rdi, (%rsp)
	cmpq	$1, (%rsp)
	jg	.LBB0_3
# %bb.1:                                # %L2
	movq	$1, 8(%rsp)
	movl	$1, %eax
	jmp	.LBB0_2
.LBB0_3:                                # %L3
	movq	(%rsp), %rbx
	leaq	-1(%rbx), %rdi
	callq	fact
	imulq	%rbx, %rax
	movq	%rax, 8(%rsp)
.LBB0_2:                                # %L2
	addq	$16, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	fact, .Lfunc_end0-fact
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L4
	pushq	%r14
	.cfi_def_cfa_offset 16
	pushq	%rbx
	.cfi_def_cfa_offset 24
	subq	$40, %rsp
	.cfi_def_cfa_offset 64
	.cfi_offset %rbx, -24
	.cfi_offset %r14, -16
	movq	$0, 8(%rsp)
	movq	$0, (%rsp)
	movq	%rsp, %r14
	leaq	16(%rsp), %rbx
	.p2align	4, 0x90
.LBB1_1:                                # %L5
                                        # =>This Inner Loop Header: Depth=1
	testb	$1, 8(%rsp)
	jne	.LBB1_4
# %bb.2:                                # %L6
                                        #   in Loop: Header=BB1_1 Depth=1
	movl	$.L.read, %edi
	movq	%r14, %rsi
	xorl	%eax, %eax
	callq	scanf
	movq	(%rsp), %rdi
	callq	fact
	movq	%rax, 24(%rsp)
	movl	$.L.fstr1, %edi
	movq	%rax, %rsi
	xorl	%eax, %eax
	callq	printf
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	cmpq	$0, 16(%rsp)
	jne	.LBB1_1
# %bb.3:                                # %L8
                                        #   in Loop: Header=BB1_1 Depth=1
	movq	$1, 8(%rsp)
	jmp	.LBB1_1
.LBB1_4:                                # %L11
	movq	$0, 32(%rsp)
	xorl	%eax, %eax
	addq	$40, %rsp
	.cfi_def_cfa_offset 24
	popq	%rbx
	.cfi_def_cfa_offset 16
	popq	%r14
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	main, .Lfunc_end1-main
	.cfi_endproc
                                        # -- End function
	.type	.L.read,@object         # @.read
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.type	.L.fstr1,@object        # @.fstr1
.L.fstr1:
	.asciz	"%ld"
	.size	.L.fstr1, 4

	.section	".note.GNU-stack","",@progbits
