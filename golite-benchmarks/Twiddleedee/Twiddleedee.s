	.text
	.file	"Twiddleedee"
	.globl	fib1                    # -- Begin function fib1
	.p2align	4, 0x90
	.type	fib1,@function
fib1:                                   # @fib1
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
	movq	(%rsp), %rax
	jmp	.LBB0_2
.LBB0_3:                                # %L3
	movq	(%rsp), %rdi
	decq	%rdi
	callq	fib1
	movq	%rax, %rbx
	movq	(%rsp), %rdi
	addq	$-2, %rdi
	callq	fib1
	addq	%rbx, %rax
.LBB0_2:                                # %L2
	movq	%rax, 8(%rsp)
	addq	$16, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	fib1, .Lfunc_end0-fib1
	.cfi_endproc
                                        # -- End function
	.globl	fib2                    # -- Begin function fib2
	.p2align	4, 0x90
	.type	fib2,@function
fib2:                                   # @fib2
	.cfi_startproc
# %bb.0:                                # %L4
	movq	%rdi, -24(%rsp)
	movq	$0, -40(%rsp)
	movq	$1, -32(%rsp)
	cmpq	$0, -24(%rsp)
	je	.LBB1_3
	.p2align	4, 0x90
.LBB1_2:                                # %L6
                                        # =>This Inner Loop Header: Depth=1
	decq	-24(%rsp)
	movq	-32(%rsp), %rax
	movq	-40(%rsp), %rcx
	addq	%rax, %rcx
	movq	%rcx, -16(%rsp)
	movq	%rax, -40(%rsp)
	movq	%rcx, -32(%rsp)
	cmpq	$0, -24(%rsp)
	jne	.LBB1_2
.LBB1_3:                                # %L7
	movq	-40(%rsp), %rax
	movq	%rax, -8(%rsp)
	retq
.Lfunc_end1:
	.size	fib2, .Lfunc_end1-fib2
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L8
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$48, %rsp
	.cfi_def_cfa_offset 64
	.cfi_offset %rbx, -16
	movl	$16, %edi
	callq	malloc
	movq	%rax, 8(%rsp)
	leaq	16(%rsp), %rbx
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	movq	16(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	movq	16(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	movq	8(%rsp), %rax
	movq	(%rax), %rdi
	callq	fib1
	movq	%rax, 32(%rsp)
	movq	8(%rsp), %rax
	movq	8(%rax), %rdi
	callq	fib2
	movq	%rax, 24(%rsp)
	movq	8(%rsp), %rdi
	callq	free
	movq	32(%rsp), %rsi
	movq	24(%rsp), %rdx
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	movq	$0, 40(%rsp)
	xorl	%eax, %eax
	addq	$48, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	main, .Lfunc_end2-main
	.cfi_endproc
                                        # -- End function
	.type	.nilnums,@object        # @.nilnums
	.comm	.nilnums,8,8
	.type	.L.read,@object         # @.read
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.type	.L.fstr1,@object        # @.fstr1
.L.fstr1:
	.asciz	"c=%ld\nd=%ld"
	.size	.L.fstr1, 12

	.section	".note.GNU-stack","",@progbits
