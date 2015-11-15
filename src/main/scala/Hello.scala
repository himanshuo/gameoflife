import models._

object Hello {
  def main(args: Array[String]): Unit = {
    val t = new Task(1,"Hi", 2)
    println(t)
  }
}
