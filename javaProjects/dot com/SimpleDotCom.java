public class SimpleDotCom {
	private int[] locationCells;
	private int numOfHits = 0;

	public void setLocationCells(int[] locs){
		locationCells = locs;
	}

	public String checkYourself(String stringGuess){
		//convert string to an int
		int guess = Integer.parseInt(stringGuess);

		/* make a variable to hold the result we'll return.
			put "miss" in as the default (assume a miss) */
		String result = "miss";
		
		//repeat with each cell in the locationCells array
		for(int cell : locationCells){
			//compare the user guess to this element, "cell" in the array
			if(guess == cell){
				result = "hit";	//neeeeeeeeeerrrrrrrr
				numOfHits++;	//booooooom
				break;
			}
		}

		//if numOfHits = 3 then dot com is dead
		if(numOfHits == locationCells.length){
			result = "kill";
		}

		//display "hit", "miss", or "kill"
		System.out.println(result);
		return result;
	}
}