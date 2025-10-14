pub enum Token {
    Constant(f64),
    Operator(Operator),
    Bracket(Side)
}

pub enum Side {Left, Right}

macro_rules! define_operators {
    ($a:ident, $b:ident, $vis:vis $enum_name:ident {$( $name:ident($operation:expr) $(,)? )*}) => {
        $vis enum $enum_name {
            $($name,)*
        }

        impl $enum_name {
            fn calculate(&self, $a: f64, $b: f64) -> f64 { match self {
                $($enum_name::$name => ($operation),)*
            }}
        }
    };
}
define_operators!(x, y, pub Operator {
    Addition(x+y),
    Negation(x-y),
    Multiplication(x*y),
    Division(x/y),
    Exponentiation(x.powf(y)),
    RootExtraction(x.powf(1f64/y))
});
