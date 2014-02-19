class GoodDog{
	//make the instance variable private
	private int size;

	//make all accessors and mutators public
	public int getSize(){ return size; }

	public void setSize(int s){ size = s; }

	void bark(){
		if(size > 60){
			System.out.println("Wooooooof");
		} else if(size > 14){
			System.out.println("Ruff");
		} else {
			System.out.println("yipyipyip");
		}
	}
}