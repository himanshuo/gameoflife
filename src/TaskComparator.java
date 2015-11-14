import java.util.Comparator;


public class TaskComparator implements Comparator<Task> {

	public int compare(Task o1, Task o2) {
		if(o1.deadline.before(o2.deadline)) return -1;
		if(o1.deadline.after(o1.deadline)) return 1;
		return 0;
		
	}
	
}
