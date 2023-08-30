	.text
	.file	"powmod"
	.globl	mod                     # -- Begin function mod
	.p2align	4, 0x90
	.type	mod,@function
mod:                                    # @mod
	.cfi_startproc
# %bb.0:                                # %L0
	movq	%rdi, -16(%rsp)
	movq	%rsi, -24(%rsp)
	movq	%rdi, %rax
	cqto
	idivq	%rsi
	movq	%rax, -32(%rsp)
	imulq	%rsi, %rax
	subq	%rax, %rdi
	movq	%rdi, -8(%rsp)
	movq	%rdi, %rax
	retq
.Lfunc_end0:
	.size	mod, .Lfunc_end0-mod
	.cfi_endproc
                                        # -- End function
	.globl	power                   # -- Begin function power
	.p2align	4, 0x90
	.type	power,@function
power:                                  # @power
	.cfi_startproc
# %bb.0:                                # %L1
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$32, %rsp
	.cfi_def_cfa_offset 48
	.cfi_offset %rbx, -16
	movq	%rdi, 8(%rsp)
	movq	%rsi, (%rsp)
	testq	%rsi, %rsi
	je	.LBB1_1
# %bb.2:                                # %L4
	movq	(%rsp), %rdi
	movl	$2, %esi
	callq	mod
	cmpq	$1, %rax
	jne	.LBB1_6
# %bb.3:                                # %L6
	movq	8(%rsp), %rbx
	movq	(%rsp), %rsi
	decq	%rsi
	movq	%rbx, %rdi
	callq	power
	imulq	%rbx, %rax
	jmp	.LBB1_4
.LBB1_1:                                # %L3
	movq	$1, 16(%rsp)
	movl	$1, %eax
	jmp	.LBB1_5
.LBB1_6:                                # %L7
	movq	8(%rsp), %rdi
	movq	(%rsp), %rax
	movq	%rax, %rsi
	shrq	$63, %rsi
	addq	%rax, %rsi
	sarq	%rsi
	callq	power
	movq	%rax, 24(%rsp)
	imulq	%rax, %rax
.LBB1_4:                                # %L6
	movq	%rax, 16(%rsp)
.LBB1_5:                                # %L6
	addq	$32, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	power, .Lfunc_end1-power
	.cfi_endproc
                                        # -- End function
	.globl	crypt                   # -- Begin function crypt
	.p2align	4, 0x90
	.type	crypt,@function
crypt:                                  # @crypt
	.cfi_startproc
# %bb.0:                                # %L8
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	movq	%rsi, 16(%rsp)
	movq	%rdx, (%rsp)
	movq	(%rdx), %rdi
	callq	power
	movq	8(%rsp), %rsi
	movq	%rax, %rdi
	callq	mod
	movq	(%rsp), %rcx
	movq	%rax, (%rcx)
	movq	(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	.nilMESSAGE(%rip), %rax
	je	.LBB2_2
# %bb.1:                                # %L10
	movq	8(%rsp), %rdi
	movq	16(%rsp), %rsi
	movq	(%rsp), %rax
	movq	8(%rax), %rdx
	callq	crypt
.LBB2_2:                                # %L12
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	crypt, .Lfunc_end2-crypt
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L13
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$80, %rsp
	.cfi_def_cfa_offset 96
	.cfi_offset %rbx, -16
	movl	$16, %edi
	callq	malloc
	movq	%rax, 24(%rsp)
	movq	%rax, 8(%rsp)
	leaq	48(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	40(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	16(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	decq	16(%rsp)
	leaq	32(%rsp), %rbx
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	movq	32(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	cmpq	$0, 16(%rsp)
	jle	.LBB3_3
	.p2align	4, 0x90
.LBB3_2:                                # %L15
                                        # =>This Inner Loop Header: Depth=1
	movl	$16, %edi
	callq	malloc
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	movq	8(%rsp), %rax
	movq	8(%rax), %rax
	movq	%rax, 8(%rsp)
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	movq	32(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, (%rcx)
	decq	16(%rsp)
	cmpq	$0, 16(%rsp)
	jg	.LBB3_2
.LBB3_3:                                # %L16
	movq	.nilMESSAGE(%rip), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	movq	40(%rsp), %rdi
	movq	48(%rsp), %rsi
	movq	24(%rsp), %rdx
	callq	crypt
	movq	24(%rsp), %rax
	movq	%rax, 8(%rsp)
	.p2align	4, 0x90
.LBB3_4:                                # %L17
                                        # =>This Inner Loop Header: Depth=1
	movq	8(%rsp), %rax
	cmpq	.nilMESSAGE(%rip), %rax
	je	.LBB3_6
# %bb.5:                                # %L18
                                        #   in Loop: Header=BB3_4 Depth=1
	movq	8(%rsp), %rax
	movq	%rax, 56(%rsp)
	movq	(%rax), %rsi
	movq	%rsi, 64(%rsp)
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	movq	8(%rsp), %rax
	movq	8(%rax), %rax
	movq	%rax, 8(%rsp)
	movq	56(%rsp), %rdi
	callq	free
	jmp	.LBB3_4
.LBB3_6:                                # %L19
	movq	$0, 72(%rsp)
	xorl	%eax, %eax
	addq	$80, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	main, .Lfunc_end3-main
	.cfi_endproc
                                        # -- End function
	.type	.nilMESSAGE,@object     # @.nilMESSAGE
	.comm	.nilMESSAGE,8,8
	.type	.L.read,@object         # @.read
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.type	.L.fstr1,@object        # @.fstr1
.L.fstr1:
	.asciz	"%ld\n"
	.size	.L.fstr1, 5

	.section	".note.GNU-stack","",@progbits
