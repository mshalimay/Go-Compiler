	.text
	.file	"primes"
	.globl	isqrt                   # -- Begin function isqrt
	.p2align	4, 0x90
	.type	isqrt,@function
isqrt:                                  # @isqrt
	.cfi_startproc
# %bb.0:                                # %L0
	movq	%rdi, -16(%rsp)
	movq	$1, -24(%rsp)
	movq	$3, -32(%rsp)
	.p2align	4, 0x90
.LBB0_1:                                # %L1
                                        # =>This Inner Loop Header: Depth=1
	movq	-24(%rsp), %rax
	cmpq	-16(%rsp), %rax
	jg	.LBB0_3
# %bb.2:                                # %L2
                                        #   in Loop: Header=BB0_1 Depth=1
	movq	-32(%rsp), %rax
	addq	%rax, -24(%rsp)
	addq	$2, %rax
	movq	%rax, -32(%rsp)
	jmp	.LBB0_1
.LBB0_3:                                # %L3
	movq	-32(%rsp), %rcx
	movq	%rcx, %rax
	shrq	$63, %rax
	addq	%rcx, %rax
	sarq	%rax
	decq	%rax
	movq	%rax, -8(%rsp)
	retq
.Lfunc_end0:
	.size	isqrt, .Lfunc_end0-isqrt
	.cfi_endproc
                                        # -- End function
	.globl	prime                   # -- Begin function prime
	.p2align	4, 0x90
	.type	prime,@function
prime:                                  # @prime
	.cfi_startproc
# %bb.0:                                # %L4
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 8(%rsp)
	cmpq	$1, %rdi
	jg	.LBB1_1
.LBB1_6:                                # %L6
	movq	$0, 16(%rsp)
	xorl	%eax, %eax
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_1:                                # %L7
	.cfi_def_cfa_offset 48
	movq	8(%rsp), %rdi
	callq	isqrt
	movq	%rax, 24(%rsp)
	movq	$2, (%rsp)
	.p2align	4, 0x90
.LBB1_2:                                # %L8
                                        # =>This Inner Loop Header: Depth=1
	movq	(%rsp), %rax
	cmpq	24(%rsp), %rax
	jg	.LBB1_5
# %bb.3:                                # %L9
                                        #   in Loop: Header=BB1_2 Depth=1
	movq	8(%rsp), %rcx
	movq	(%rsp), %rsi
	movq	%rcx, %rax
	cqto
	idivq	%rsi
	imulq	%rsi, %rax
	subq	%rax, %rcx
	movq	%rcx, 32(%rsp)
	je	.LBB1_6
# %bb.4:                                # %L12
                                        #   in Loop: Header=BB1_2 Depth=1
	incq	(%rsp)
	jmp	.LBB1_2
.LBB1_5:                                # %L14
	movq	$1, 16(%rsp)
	movl	$1, %eax
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	prime, .Lfunc_end1-prime
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L15
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	leaq	8(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movq	$0, (%rsp)
	jmp	.LBB2_1
	.p2align	4, 0x90
.LBB2_4:                                # %L21
                                        #   in Loop: Header=BB2_1 Depth=1
	incq	(%rsp)
.LBB2_1:                                # %L16
                                        # =>This Inner Loop Header: Depth=1
	movq	(%rsp), %rax
	cmpq	8(%rsp), %rax
	jg	.LBB2_5
# %bb.2:                                # %L17
                                        #   in Loop: Header=BB2_1 Depth=1
	movq	(%rsp), %rdi
	callq	prime
	testb	$1, %al
	je	.LBB2_4
# %bb.3:                                # %L19
                                        #   in Loop: Header=BB2_1 Depth=1
	movq	(%rsp), %rsi
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	jmp	.LBB2_4
.LBB2_5:                                # %L22
	movq	$0, 16(%rsp)
	xorl	%eax, %eax
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	main, .Lfunc_end2-main
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
