
//Random comment
public class MyBean {
    private String first, second, third = "value";
    private String someProperty;
    public int someProperty2;
    protected String someProperty3;
    int someInt;

    @Annotated
    public String getSomeProperty() {
        return someProperty;
    }

    public void setSomeProperty(String someProperty) {
        this.someProperty = someProperty;
    }

    public static void someStatic() {
        return null;
    }
}