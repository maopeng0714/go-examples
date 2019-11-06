package visitTree;

import java.util.ArrayList;
import java.util.LinkedList;


public class visitTree {
	public static class Node {
		public Node left;
		public Node right;
		public int value;
		public Node(int value){
			this.value = value;
		}
		public Node getLeft(){
			return left;
		}
		public void setLeft(Node left){
			this.left = left;
		}
		public Node getRight(){
			return right;
		}
		public void setRight(Node right){
			this.right = right;
		}
	}
	public visitTree(){

	}
	
	public void visit(Node root) {
		 LinkedList queue = new LinkedList<Node>();
		
		queue.offer(root);
		while (!queue.isEmpty()) {
			Node ele = (Node)queue.poll();
			System.out.println(ele.value);
			if(ele.getLeft() != null){
				queue.offer(ele.getLeft());
			}
			if(ele.getRight() != null){
				queue.offer(ele.getRight());
			}
		}
	   
	}
	public static void main(String[] args){
		visitTree visitInstance = new visitTree();

		Node root = new Node(1);
		root.setLeft( new Node(2));
		root.setRight( new Node(3));
		root.left.setLeft(new Node(4));
		root.left.setRight(new Node(5));
		root.right.setLeft(new Node(6));
		root.right.setRight(new Node(7));
		visitInstance.visit(root);
	}
}

