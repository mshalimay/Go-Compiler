	.text
	.file	"linkedlist"
	.globl	Add                     # -- Begin function Add
	.p2align	4, 0x90
	.type	Add,@function
Add:                                    # @Add
	.cfi_startproc
# %bb.0:                                # %L0
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 16(%rsp)
	movl	$16, %edi
	callq	malloc
	movq	%rax, 8(%rsp)
	movq	16(%rsp), %rcx
	movq	%rcx, (%rax)
	movq	.nilNode(%rip), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	movq	head(%rip), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB0_1
# %bb.2:                                # %L3
	movq	8(%rsp), %rax
	movq	tail(%rip), %rcx
	movq	%rax, 8(%rcx)
	movq	8(%rsp), %rax
	jmp	.LBB0_3
.LBB0_1:                                # %L2
	movq	8(%rsp), %rax
	movq	%rax, head(%rip)
.LBB0_3:                                # %L4
	movq	%rax, tail(%rip)
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end0:
	.size	Add, .Lfunc_end0-Add
	.cfi_endproc
                                        # -- End function
	.globl	PrintList               # -- Begin function PrintList
	.p2align	4, 0x90
	.type	PrintList,@function
PrintList:                              # @PrintList
	.cfi_startproc
# %bb.0:                                # %L5
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	cmpq	tail(%rip), %rdi
	je	.LBB1_1
# %bb.2:                                # %L8
	movq	8(%rsp), %rax
	movq	(%rax), %rsi
	movq	%rsi, 16(%rsp)
	movl	$.L.fstr2, %edi
	xorl	%eax, %eax
	callq	printf
	movq	8(%rsp), %rax
	movq	8(%rax), %rdi
	callq	PrintList
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_1:                                # %L7
	.cfi_def_cfa_offset 32
	movq	8(%rsp), %rax
	movq	(%rax), %rsi
	movq	%rsi, 16(%rsp)
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	PrintList, .Lfunc_end1-PrintList
	.cfi_endproc
                                        # -- End function
	.globl	Del                     # -- Begin function Del
	.p2align	4, 0x90
	.type	Del,@function
Del:                                    # @Del
	.cfi_startproc
# %bb.0:                                # %L10
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, (%rsp)
	movq	%rsi, 16(%rsp)
	cmpq	.nilNode(%rip), %rdi
	je	.LBB2_10
# %bb.1:                                # %L13
	movq	head(%rip), %rax
	movq	(%rax), %rax
	cmpq	16(%rsp), %rax
	jne	.LBB2_3
# %bb.2:                                # %L15
	movq	head(%rip), %rdi
	movq	%rdi, 8(%rsp)
	movq	8(%rdi), %rax
	movq	%rax, head(%rip)
	jmp	.LBB2_8
.LBB2_3:                                # %L16
	movq	(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	tail(%rip), %rax
	je	.LBB2_6
# %bb.4:                                # %L19
	movq	(%rsp), %rax
	movq	8(%rax), %rax
	movq	(%rax), %rax
	cmpq	16(%rsp), %rax
	jne	.LBB2_9
# %bb.5:                                # %L21
	movq	(%rsp), %rax
	movq	8(%rax), %rcx
	movq	%rcx, 8(%rsp)
	movq	8(%rax), %rcx
	movq	8(%rcx), %rcx
	jmp	.LBB2_7
.LBB2_6:                                # %L18
	movq	tail(%rip), %rax
	movq	%rax, 8(%rsp)
	movq	(%rsp), %rax
	movq	%rax, tail(%rip)
	movq	.nilNode(%rip), %rcx
.LBB2_7:                                # %L26
	movq	%rcx, 8(%rax)
	movq	8(%rsp), %rdi
.LBB2_8:                                # %L26
	callq	free
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB2_9:                                # %L22
	.cfi_def_cfa_offset 32
	movq	(%rsp), %rax
	movq	8(%rax), %rdi
	movq	16(%rsp), %rsi
	callq	Del
.LBB2_10:                               # %L26
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	Del, .Lfunc_end2-Del
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L27
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	leaq	24(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	leaq	16(%rsp), %rsi
	movl	$.L.read, %edi
	xorl	%eax, %eax
	callq	scanf
	movl	$1, %edi
	callq	Add
	movl	$10, %edi
	callq	Add
	movl	$3, %edi
	callq	Add
	movl	$4, %edi
	callq	Add
	movq	24(%rsp), %rdi
	callq	Add
	movq	head(%rip), %rdi
	callq	PrintList
	movq	$0, 8(%rsp)
	cmpq	$49999999, 8(%rsp)      # imm = 0x2FAF07F
	jg	.LBB3_3
	.p2align	4, 0x90
.LBB3_2:                                # %L29
                                        # =>This Inner Loop Header: Depth=1
	movq	8(%rsp), %rdi
	callq	Add
	incq	8(%rsp)
	cmpq	$49999999, 8(%rsp)      # imm = 0x2FAF07F
	jle	.LBB3_2
.LBB3_3:                                # %L30
	movq	$0, 8(%rsp)
	cmpq	$49999999, 8(%rsp)      # imm = 0x2FAF07F
	jg	.LBB3_6
	.p2align	4, 0x90
.LBB3_5:                                # %L32
                                        # =>This Inner Loop Header: Depth=1
	movq	head(%rip), %rdi
	movq	8(%rsp), %rsi
	callq	Del
	incq	8(%rsp)
	cmpq	$49999999, 8(%rsp)      # imm = 0x2FAF07F
	jle	.LBB3_5
.LBB3_6:                                # %L33
	movq	head(%rip), %rdi
	movq	16(%rsp), %rsi
	callq	Del
	movq	head(%rip), %rdi
	callq	PrintList
	movq	$0, 32(%rsp)
	xorl	%eax, %eax
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	main, .Lfunc_end3-main
	.cfi_endproc
                                        # -- End function
	.type	.nilNode,@object        # @.nilNode
	.comm	.nilNode,8,8
	.type	.nilNode2,@object       # @.nilNode2
	.comm	.nilNode2,8,8
	.type	head,@object            # @head
	.comm	head,8,8
	.type	tail,@object            # @tail
	.comm	tail,8,8
	.type	.L.fstr1,@object        # @.fstr1
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.fstr1:
	.asciz	"%ld"
	.size	.L.fstr1, 4

	.type	.L.fstr2,@object        # @.fstr2
.L.fstr2:
	.asciz	"%ld"
	.size	.L.fstr2, 4

	.type	.L.read,@object         # @.read
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.section	".note.GNU-stack","",@progbits
