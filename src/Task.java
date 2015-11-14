import java.util.Date;
import java.util.List;
import java.util.Scanner;


public class Task {
	
	
	public int id;	
	public String name;
	public String description;
	public String acceptanceCriteria;
	public String failureCriteria;
	public Date deadline;
	public int points;
	public List<Task> subtasks;
	public String goal;
	public Progress progress;
	public boolean recurring;
	
	
	
	public boolean finish(){
		
		Scanner scanner = new Scanner(System.in);
		System.out.printf("Acceptance Criteria: %s\n", this.acceptanceCriteria);
		System.out.printf("Failure Criteria: %s\n", this.failureCriteria);
		System.out.print("Can we accept, fail, or neither (a/f/n):");
		String result = scanner.nextLine();
		if(result.equalsIgnoreCase("a")){
			//make sure all subtasks are done
			for(Task s: this.subtasks){
				if(s.progress != Progress.DONE) return false;
			}
			this.progress = Progress.DONE;
			return true;
		} else {
			return false;
		}
		
	}
	
	public String toString(){
		//todo: format the date properly
		return String.format("<Task (%d) %s :  DUE:%s>", this.id,this.name, this.deadline);
	}
	
	public boolean valid(){
		//todo: handle still in appropriate time
		return this.progress != Progress.DONE;
	}
}
