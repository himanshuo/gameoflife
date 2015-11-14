import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.List;

/**
 * @author himanshu
 *
 */
public class Player {
	public String name;
	public int achievementPoints;
	public int himanshuPoints;
	public List<FunCard> funCards;
	public List<Task> tasks;
	public List<Goal> goals;
	
	public Player(String name){
		this.name=name;
		this.achievementPoints=0;
		this.himanshuPoints=0;
		this.funCards = new ArrayList<FunCard>();
		this.tasks = new ArrayList<Task>();
		this.goals = new ArrayList<Goal>();
	}
	
	
	public FunCard drawFunCard(){
		List<FunCard> validFunCards = new ArrayList<FunCard>();
		for(FunCard fc: this.funCards){
			if(fc.valid()){
				validFunCards.add(fc);
			}
		}
		if(validFunCards.size()==0){
			//todo: exceptions
			//throw new NoFunCardsException();
		}
			
		Collections.shuffle(validFunCards);
		return validFunCards.get(0);
	}
	
	//todo: move into Game
	public List<Task> openTasks(){
		List<Task> validTasks = new ArrayList<Task>();
		for(Task t: this.tasks){
			if(t.valid()){
				validTasks.add(t);
			}
		}
		
		Collections.sort(validTasks, new TaskComparator());
		return validTasks;
	}
	
	public boolean finishTask(Task t){
		if(t.finish()){
			//todo: make sure you aren't past the deadline
			
			this.achievementPoints += t.points;
			return true;
		} else{
			this.achievementPoints -= t.points;
			return false;
		}
		
	}
	
	public boolean updateTask(Task t, Progress p){
		t.progress = p;
		return true;
	}
	
}
