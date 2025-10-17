use derive_more::Display;

#[derive(Display)]
#[display("{_0}")]
pub enum Token {
    Constant(f64),
    Operator(Operator),
    Bracket(Side)
}

#[derive(Display)]
pub enum Side {
    #[display("(")]
    Left,
    #[display(")")]
    Right
}


macro_rules! define_operators {
    ($a:ident, $b:ident, $vis:vis $enum_name:ident {$( $name:ident($identifier:expr) = $operation:expr ),*}) => {
        #[derive(Display)]
        $vis enum $enum_name {
            $(
                #[display($identifier)]
                $name
            ),*
        }

        impl $enum_name {
            fn calculate(&self, $a: f64, $b: f64) -> f64 { match self {
                $($enum_name::$name => ($operation),)*
            }}

            fn identify(operator_str: &str) -> Option<Operator> { match operator_str {
                $( $identifier => Some(Operator::$name),)*
                _ => None
            }}
        }
    };
}
define_operators!(x, y, pub Operator {
    Addition("+") = x+y,
    Negation("-") = x-y,
    Multiplication("*") = x*y,
    Division("/") = x/y,
    Exponentiation("^") = x.powf(y),
    RootExtraction("√") = x.powf(1f64/y)
});
