import java.util.Scanner;


public class Game {
	
	public Player player;
	public Game(){
		this.player = new Player("Himanshu");
	}
	
	public void showScoreBoard(){
		System.out.printf("Player %s\nHimanshu Points: %d, Achiement Points: %d", 
				this.player.name,
				this.player.achievementPoints,
				this.player.himanshuPoints
				);
		this.showOpenTasks();
		
	}
	
	public void showOpenTasks(){
		for(Task t: this.player.openTasks()){
			System.out.println(t);
		}
	}
	
	public void viewOpenTasksAction(){
		System.out.println("Open tasks (tasks that are not done or expired):");
		this.showOpenTasks();
		System.out.println("Choose an option to update the ");
		
	}
	
	public static void main(String args[]){
		Game game = new Game();
		boolean playing = true; 
		Scanner scanner = new Scanner(System.in);
		while(playing){
			game.showScoreBoard();
			String options[] = {
					"Quit", 
					"View Open Tasks", 
					"Draw Fun Card", 
					"View Specific Task",
					"Update Task Progress",
					"Update Task Details",
					""
					};
			for(int i=0;i<options.length;i++){
				System.out.printf("%d: %s\n",i, options[i]);
			}			
			System.out.print("Choose an option:");
			int n = scanner.nextInt();
			switch(n){
			case 0: playing=false; break;
			case 1: game.showOpenTasks(); break;
			case 2: game.player.drawFunCard(); break;
			case 3: 
			}
			
		}
		
		
	}

}
