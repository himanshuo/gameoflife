import java.util.Date;


public class FunCard {
	public String description;
	public Date duration;
	public boolean used;
	public int points;
	
	public FunCard(String desc, Date dur, int points){
		this.description = desc;
		this.duration = dur;
		this.used=false;
		this.points = points;
	}
	
	
	public boolean use(Player p){
		//todo: check evennt hasnt exired
		
		p.achievementPoints += this.points;
		return true;
	}
	
	
	public boolean valid(){
		//todo: check if endtime is before current time.
		return !this.used;
	}
}
