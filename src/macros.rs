#[macro_export]
macro_rules! define_operators {
    ($a:ident, $b:ident, $vis:vis $enum_name:ident {$( $name:ident($identifier:expr) = $operation:expr ),*}) => {
        #[derive(Display, Debug)]
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
