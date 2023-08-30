	.text
	.file	"binarytree"
	.globl	compare                 # -- Begin function compare
	.p2align	4, 0x90
	.type	compare,@function
compare:                                # @compare
	.cfi_startproc
# %bb.0:                                # %L0
	movq	%rdi, -8(%rsp)
	movq	%rsi, -16(%rsp)
	cmpq	%rsi, %rdi
	jge	.LBB0_1
# %bb.3:                                # %L2
	movq	$1, -24(%rsp)
	movl	$1, %eax
	retq
.LBB0_1:                                # %L3
	movq	-8(%rsp), %rax
	cmpq	-16(%rsp), %rax
	jle	.LBB0_2
# %bb.4:                                # %L5
	movq	$-1, -24(%rsp)
	movq	$-1, %rax
	retq
.LBB0_2:                                # %L6
	movq	$0, -24(%rsp)
	xorl	%eax, %eax
	retq
.Lfunc_end0:
	.size	compare, .Lfunc_end0-compare
	.cfi_endproc
                                        # -- End function
	.globl	addNode                 # -- Begin function addNode
	.p2align	4, 0x90
	.type	addNode,@function
addNode:                                # @addNode
	.cfi_startproc
# %bb.0:                                # %L7
	subq	$40, %rsp
	.cfi_def_cfa_offset 48
	movq	%rdi, 16(%rsp)
	movq	%rsi, 8(%rsp)
	cmpq	.nilNode(%rip), %rsi
	je	.LBB1_1
# %bb.2:                                # %L10
	movq	8(%rsp), %rax
	movq	(%rax), %rdi
	movq	16(%rsp), %rsi
	callq	compare
	movq	%rax, 32(%rsp)
	cmpq	$-1, %rax
	jne	.LBB1_6
# %bb.3:                                # %L12
	movq	8(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB1_4
# %bb.5:                                # %L16
	movq	16(%rsp), %rdi
	movq	8(%rsp), %rax
	movq	8(%rax), %rsi
	jmp	.LBB1_10
.LBB1_1:                                # %L9
	movl	$24, %edi
	callq	malloc
	movq	%rax, 24(%rsp)
	movq	16(%rsp), %rcx
	movq	%rcx, (%rax)
	movq	24(%rsp), %rax
	movq	%rax, root(%rip)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_6:                                # %L13
	.cfi_def_cfa_offset 48
	cmpq	$1, 32(%rsp)
	jne	.LBB1_11
# %bb.7:                                # %L19
	movq	8(%rsp), %rax
	movq	16(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB1_8
# %bb.9:                                # %L23
	movq	16(%rsp), %rdi
	movq	8(%rsp), %rax
	movq	16(%rax), %rsi
.LBB1_10:                               # %L27
	callq	addNode
.LBB1_11:                               # %L27
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_4:                                # %L15
	.cfi_def_cfa_offset 48
	movl	$24, %edi
	callq	malloc
	movq	%rax, 24(%rsp)
	movq	16(%rsp), %rcx
	movq	%rcx, (%rax)
	movq	24(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 8(%rcx)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.LBB1_8:                                # %L22
	.cfi_def_cfa_offset 48
	movl	$24, %edi
	callq	malloc
	movq	%rax, 24(%rsp)
	movq	16(%rsp), %rcx
	movq	%rcx, (%rax)
	movq	24(%rsp), %rax
	movq	8(%rsp), %rcx
	movq	%rax, 16(%rcx)
	addq	$40, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end1:
	.size	addNode, .Lfunc_end1-addNode
	.cfi_endproc
                                        # -- End function
	.globl	printDepthTree          # -- Begin function printDepthTree
	.p2align	4, 0x90
	.type	printDepthTree,@function
printDepthTree:                         # @printDepthTree
	.cfi_startproc
# %bb.0:                                # %L28
	subq	$24, %rsp
	.cfi_def_cfa_offset 32
	movq	%rdi, 8(%rsp)
	cmpq	.nilNode(%rip), %rdi
	je	.LBB2_5
# %bb.1:                                # %L30
	movq	8(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB2_3
# %bb.2:                                # %L33
	movq	8(%rsp), %rax
	movq	8(%rax), %rdi
	callq	printDepthTree
.LBB2_3:                                # %L35
	movq	8(%rsp), %rax
	movq	(%rax), %rsi
	movq	%rsi, 16(%rsp)
	movl	$.L.fstr1, %edi
	xorl	%eax, %eax
	callq	printf
	movq	8(%rsp), %rax
	movq	16(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB2_5
# %bb.4:                                # %L37
	movq	8(%rsp), %rax
	movq	16(%rax), %rdi
	callq	printDepthTree
.LBB2_5:                                # %L40
	addq	$24, %rsp
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end2:
	.size	printDepthTree, .Lfunc_end2-printDepthTree
	.cfi_endproc
                                        # -- End function
	.globl	deleteLeavesTree        # -- Begin function deleteLeavesTree
	.p2align	4, 0x90
	.type	deleteLeavesTree,@function
deleteLeavesTree:                       # @deleteLeavesTree
	.cfi_startproc
# %bb.0:                                # %L41
	pushq	%rax
	.cfi_def_cfa_offset 16
	movq	%rdi, (%rsp)
	cmpq	.nilNode(%rip), %rdi
	je	.LBB3_6
# %bb.1:                                # %L43
	movq	(%rsp), %rax
	movq	8(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB3_3
# %bb.2:                                # %L46
	movq	(%rsp), %rax
	movq	8(%rax), %rdi
	callq	deleteLeavesTree
.LBB3_3:                                # %L48
	movq	(%rsp), %rax
	movq	16(%rax), %rax
	cmpq	.nilNode(%rip), %rax
	je	.LBB3_5
# %bb.4:                                # %L50
	movq	(%rsp), %rax
	movq	16(%rax), %rdi
	callq	deleteLeavesTree
.LBB3_5:                                # %L52
	movq	(%rsp), %rdi
	callq	free
.LBB3_6:                                # %L53
	popq	%rax
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end3:
	.size	deleteLeavesTree, .Lfunc_end3-deleteLeavesTree
	.cfi_endproc
                                        # -- End function
	.globl	main                    # -- Begin function main
	.p2align	4, 0x90
	.type	main,@function
main:                                   # @main
	.cfi_startproc
# %bb.0:                                # %L54
	pushq	%rbx
	.cfi_def_cfa_offset 16
	subq	$32, %rsp
	.cfi_def_cfa_offset 48
	.cfi_offset %rbx, -16
	movq	.nilNode(%rip), %rax
	movq	%rax, root(%rip)
	movq	$0, 8(%rsp)
	leaq	8(%rsp), %rbx
	.p2align	4, 0x90
.LBB4_1:                                # %L55
                                        # =>This Inner Loop Header: Depth=1
	movl	$.L.read, %edi
	movq	%rbx, %rsi
	xorl	%eax, %eax
	callq	scanf
	cmpq	$0, 8(%rsp)
	je	.LBB4_3
# %bb.2:                                # %L56
                                        #   in Loop: Header=BB4_1 Depth=1
	movq	8(%rsp), %rdi
	movq	root(%rip), %rsi
	callq	addNode
	jmp	.LBB4_1
.LBB4_3:                                # %L57
	movq	root(%rip), %rdi
	callq	printDepthTree
	movq	root(%rip), %rdi
	callq	deleteLeavesTree
	movq	$0, 16(%rsp)
	xorl	%eax, %eax
	addq	$32, %rsp
	.cfi_def_cfa_offset 16
	popq	%rbx
	.cfi_def_cfa_offset 8
	retq
.Lfunc_end4:
	.size	main, .Lfunc_end4-main
	.cfi_endproc
                                        # -- End function
	.type	.nilNode,@object        # @.nilNode
	.comm	.nilNode,8,8
	.type	root,@object            # @root
	.comm	root,8,8
	.type	.L.fstr1,@object        # @.fstr1
	.section	.rodata.str1.1,"aMS",@progbits,1
.L.fstr1:
	.asciz	"%ld"
	.size	.L.fstr1, 4

	.type	.L.read,@object         # @.read
.L.read:
	.asciz	"%ld"
	.size	.L.read, 4

	.section	".note.GNU-stack","",@progbits
